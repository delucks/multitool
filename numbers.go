package main

import (
	"fmt"
	"io"
	"strconv"
)

func ConvertBase(args []string, _ io.Reader) error {
	if len(args) != 3 {
		return fmt.Errorf("basejump: two arguments are required\n  basejump <input integer> <output base>\n  basejump 0xfeedface 10\n  basejump 258 2\n  basejump 0b1001011 10\n  basejump 0o23815 2")
	}
	// Parse the output base
	output_base, err := strconv.ParseUint(args[2], 10, 64)
	if err != nil {
		return err
	}
	if output_base < 2 || output_base > 36 {
		return fmt.Errorf("The output base cannot be less than 2 or higher than 36.")
	}
	// The first two characters are either a base specification or the start of the number. Let's find out.
	var arg string
	var input_base int
	//fmt.Printf("%v\n", args[1][:2])
	switch args[1][:2] {
	case "0x", "0X":
		input_base = 16
		arg = args[1][2:]
	case "0b", "0B":
		input_base = 2
		arg = args[1][2:]
	case "0o", "0O":
		input_base = 8
		arg = args[1][2:]
	default:
		// No base specification provided
		input_base = 10
		arg = args[1]
	}
	i64, err := strconv.ParseInt(arg, input_base, 64)
	if err != nil {
		return err
	}
	fmt.Println(strconv.FormatInt(i64, int(output_base)))
	return nil
}
