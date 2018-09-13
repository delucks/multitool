package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var useANSIColor bool

// Outputs messages with ISO8601 timestamps and terminal colors for severity levels
func ShellLogger(args []string, stdin io.Reader) error {
	var severity, message string
	args = args[1:] // We don't need "log"
	switch len(args) {
	case 1:
		severity = "info"
		message = args[0]
	case 2:
		message = args[0]
		severity = args[1]
	default:
		return fmt.Errorf("log needs at least one argument, the message to log. If a second argument is provided, it's the message's severity. info|warning|error|debug have different colors.")
	}
	severity_map := map[string]ANSIColor{
		"err":       Red,
		"error":     Red,
		"exception": Red,
		"fatal":     Red,
		"warn":      Yellow,
		"warning":   Yellow,
		"ok":        Green,
		"info":      Green,
		"debug":     Blue,
	}
	// Should we use ANSI terminal sequences to get colors?
	switch os.Getenv("COLOR_ENABLED") {
	case "false", "no", "FALSE", "f", "n":
		useANSIColor = false
	default:
		useANSIColor = true
	}
	now := time.Now().Format(time.RFC3339) + "\t"
	if useANSIColor {
		color, ok := severity_map[severity]
		if !ok {
			// Default
			color = Green
		}
		fmt.Fprintln(os.Stderr, ColorWrap(now+message, color))
		return nil
	}
	fmt.Fprintln(os.Stderr, now+message)
	return nil
}
