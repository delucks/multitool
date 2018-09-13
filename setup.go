package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SymlinkAll(_ []string, _ io.Reader) error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	bindir := filepath.Dir(self)
	fmt.Fprintf(os.Stderr, "Symlinking tools into %s\n", bindir)
	for k := range Entrypoints {
		switch k {
		case "help", "setup":
			continue
		default:
			dest := bindir + "/" + k
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
