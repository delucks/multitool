package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Tool struct {
	Usage      string
	Entrypoint func([]string, io.Reader) error
}

var Entrypoints map[string]Tool

func Help([]string, io.Reader) error {
	fmt.Fprintf(os.Stderr, "multitool: https://github.com/delucks/multitool\n\n")
	fmt.Fprintf(os.Stderr, "%-15sDescription\n", "Tool")
	fmt.Fprintf(os.Stderr, "%-15s-----------\n", "----")
	for k, v := range Entrypoints {
		fmt.Fprintf(os.Stderr, "%-15s%s\n", k, v.Usage)
	}
	return nil
}

func main() {
	Entrypoints = map[string]Tool{
		"basejump":   Tool{"Convert an integer between base representations", ConvertBase},
		"log":        Tool{"Simple logger for use in shell scripts", ShellLogger},
		"suggest-fc": Tool{"Spell-correct an incorrectly typed executable", SpellCorrectCommand},
		"help":       Tool{"Display this help output", Help},
	}

	which := filepath.Base(os.Args[0])
	args := os.Args
	if which == "multitool" && len(args) > 1 {
		args = args[1:]
		which = args[0]
	}

	fn, ok := Entrypoints[which]
	if !ok {
		Help(args, os.Stdin)
		os.Exit(1)
	}
	err := fn.Entrypoint(args, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
