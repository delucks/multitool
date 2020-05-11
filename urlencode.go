package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
)

func UrlEncode(args []string, stdin io.Reader) error {
	if len(args) >= 2 {
		for _, arg := range args[1:] {
			fmt.Println(url.QueryEscape(arg))
		}
	}
	stat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			currentLine := scanner.Text()
			fmt.Println(url.QueryEscape(currentLine))
		}
		if err = scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}
