package main

import (
	"fmt"
	"os"
	"strings"
)

func ls(c *CLI, args []string) int {
	if len(args) > 3 {
		mArgsPrint(c)
		return ExitCodeError
	}

	var path string
	// When path argument is given
	if len(args) == 3 {
		path = args[2]
	}
	// When no argument
	if len(args) == 2 {
		currentDir, err := os.Getwd()
		if err != nil {
			errPrint(c, err)
			return ExitCodeError
		}
		path = currentDir
	}

	files, err := os.ReadDir(path)
	if err != nil {
		errPrint(c, err)
		return ExitCodeError
	}

	for _, file := range files {
		// skip hidden file
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		fmt.Fprintf(c.outStream, "%s\n", file.Name())
	}
	return ExitCodeOk
}
