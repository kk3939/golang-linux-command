package main

import (
	"fmt"
	"os"
)

func cat(c *CLI, args []string) int {
	// When too much arguments
	if len(args) > 3 {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "expected arguments are too much!")
		return ExitCodeError
	}

	// When no argument
	if len(args) == 2 {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "argument is required!")
		return ExitCodeError
	}

	file, err := os.ReadFile(args[2])
	if err != nil {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", err)
		return ExitCodeError
	}
	fmt.Fprintf(c.outStream, "%s", string(file))
	return ExitCodeOk
}
