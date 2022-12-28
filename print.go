package main

import "fmt"

func errPrint(c *CLI, err error) {
	fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", err)
}

func mArgsPrint(c *CLI) {
	fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "expected arguments are too much!")
}

func rArgsPrint(c *CLI) {
	fmt.Fprintf(c.errStream, "golang-linux-command: %s\n", "argument is required!")
}
