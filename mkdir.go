package main

import (
	"io/fs"
	"os"
)

var createDir func(name string, perm fs.FileMode) error

func init() {
	createDir = os.Mkdir
}

func SetCreateDir(f func(name string, perm fs.FileMode) error) func() {
	// https://medium.com/veltra-engineering/how-to-write-test-code-for-file-output-in-golang-54e200872e8a
	tempF := createDir
	createDir = f
	return func() {
		createDir = tempF
	}
}

func mkdir(c *CLI, args []string) int {
	if len(args) > 3 {
		mArgsPrint(c)
		return ExitCodeError
	}

	if len(args) == 2 {
		rArgsPrint(c)
		return ExitCodeError
	}

	currentDir, err := os.Getwd()
	if err != nil {
		errPrint(c, err)
		return ExitCodeError
	}

	fileInfo, err := os.Lstat(currentDir)
	if err != nil {
		errPrint(c, err)
	}

	// How to specify permission of directory
	// https://twitter.com/mattn_jp/status/1546039502439055363
	if err := createDir(args[2], fileInfo.Mode()&os.ModePerm); err != nil {
		return ExitCodeError
	}
	return ExitCodeOk
}
