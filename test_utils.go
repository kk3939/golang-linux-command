package main

import (
	"bytes"
	"strings"
)

func setup_test(a string) (*CLI, []string, *bytes.Buffer, *bytes.Buffer) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split(a, " ")
	return cli, args, outStream, errStream
}
