package main

import (
	"fmt"
	"os"
)

func mkdir(c *CLI, args []string) int {
	if len(args) > 3 {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "expected arguments are too much!")
		return ExitCodeError
	}

	if len(args) == 2 {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "argument is required!")
		return ExitCodeError
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", err)
		return ExitCodeError
	}

	fileInfo, err := os.Lstat(currentDir)
	if err != nil {
		fmt.Fprintf(c.outStream, "golang-linux-command: %s\n", err)
	}

	// How to specify permission of directory
	// https://twitter.com/mattn_jp/status/1546039502439055363
	if err := os.Mkdir(args[2], fileInfo.Mode()&os.ModePerm); err != nil {
		return ExitCodeError
	}
	return ExitCodeOk
}
