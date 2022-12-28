package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_ls(t *testing.T) {
	t.Run("no argument", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := []string{"golang-linux-command", "ls"}

		if s := ls(cli, args); s != ExitCodeOk {
			t.Errorf("Exit code is %v", s)
		}
	})

	t.Run("one argument", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := []string{"golang-linux-command", "ls", "testdata"}

		if s := ls(cli, args); s != ExitCodeOk {
			t.Errorf("Exit code is %v", s)
		} else {
			if !strings.Contains(outStream.String(), "test1.txt\ntest2.txt") {
				t.Errorf("Files are contained in outStream.")
			}
		}
	})

	t.Run("argument_is_not_dir", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := []string{"golang-linux-command", "ls", "README.md"}

		if s := ls(cli, args); s == ExitCodeOk {
			t.Errorf("Exit code is %v", s)
		}
	})

	t.Run("one more argument", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := []string{"golang-linux-command", "ls", "testdata", "testdata"}

		if s := ls(cli, args); s == ExitCodeOk {
			t.Errorf("it is Expected that it would be failed. Exit code is %v. Too much arguments", s)
		}
	})
}
