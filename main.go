package main

import (
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeZero = 0
	ExitCodeOne  = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	if len(args) == 1 {
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "specify linux command to a first argument!")
		return ExitCodeOne
	}

	switch args[1] {
	case "ls":
		return ls(c, args)
	case "cat":
		return cat(c, args)
	default:
		fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "command is not implemented!")
		return ExitCodeOne
	}
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	status := cli.Run(os.Args)
	os.Exit(status)
}
