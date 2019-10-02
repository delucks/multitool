package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func pathIsDir(path string) (bool, error) {
	fh, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fh.IsDir(), nil
}

func SymlinkAll(args []string, _ io.Reader) error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	destinationDir := filepath.Dir(self)
	if len(args) >= 2 {
		// User may have passed an alternate destination directory
		bindir := args[1]
		ok, err := pathIsDir(bindir)
		if err != nil {
			return err
		}
		if ok {
			destinationDir = bindir
		} else {
			return errors.New(bindir + " is not a valid directory")
		}
	}
	fmt.Fprintf(os.Stderr, "Symlinking tools into %s\n", destinationDir)
	for k := range Entrypoints {
		switch k {
		case "help", "setup":
			continue
		default:
			dest := destinationDir + "/" + k
			// Check to see if the path already exists
			if _, err = os.Stat(dest); err == nil {
				resolved, err := filepath.EvalSymlinks(dest)
				if err != nil {
					return err
				}
				rel, err := filepath.Rel(self, resolved)
				if rel != "." {
					// This symlink points to a location other than the current binary!
					if err = os.Remove(dest); err != nil {
						return err
					}
					goto Symlink
				}
				// The path points back to our binary already, we don't have to do anything
				fmt.Fprintf(os.Stderr, "%s exists already\n", k)
				continue
			}
		Symlink:
			err = os.Symlink(self, dest)
			if err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "%s symlinked\n", k)
		}
	}
	return nil
}
