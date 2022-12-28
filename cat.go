package main

import (
	"fmt"
	"os"
)

func cat(c *CLI, args []string) int {
	// When too much arguments
	if len(args) > 3 {
		mArgsPrint(c)
		return ExitCodeError
	}

	// When no argument
	if len(args) == 2 {
		rArgsPrint(c)
		return ExitCodeError
	}

	file, err := os.ReadFile(args[2])
	if err != nil {
		errPrint(c, err)
		return ExitCodeError
	}
	fmt.Fprintf(c.outStream, "%s", string(file))
	return ExitCodeOk
}
