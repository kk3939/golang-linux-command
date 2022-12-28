package main

import (
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOk    = 0
	ExitCodeError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	if len(args) == 1 {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "specify linux command to a first argument!")
		return ExitCodeError
	}

	switch args[1] {
	case "ls":
		return ls(c, args)
	case "cat":
		return cat(c, args)
	case "mkdir":
		return mkdir(c, args)
	default:
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "command is not implemented!")
		return ExitCodeError
	}
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	status := cli.Run(os.Args)
	os.Exit(status)
}
