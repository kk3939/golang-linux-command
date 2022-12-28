package main

import (
	"strings"
	"testing"
)

func Test_cat(t *testing.T) {
	t.Run("no_argument", func(t *testing.T) {
		cli, args, _, _ := setup_test("golang-linux-command cat")
		if s := cat(cli, args); s == ExitCodeOk {
			t.Errorf("Exit code is %v", s)
		}
	})

	t.Run("one_argument", func(t *testing.T) {
		cli, args, outStream, _ := setup_test("golang-linux-command cat README.md")
		if s := cat(cli, args); s == ExitCodeError {
			t.Errorf("Exit code is %v", s)
		} else {
			if !strings.Contains(outStream.String(), "golang-linux-command") {
				t.Errorf("String of h1 in README.md are contained in outStream.")
			}
		}
	})

	t.Run("argument_is_dir", func(t *testing.T) {
		cli, args, _, _ := setup_test("golang-linux-command cat testdata")
		if s := cat(cli, args); s == ExitCodeOk {
			t.Errorf("Exit code is %v", s)
		}
	})

	t.Run("one more argument", func(t *testing.T) {
		cli, args, _, _ := setup_test("golang-linux-command cat README.md testdata")
		if s := cat(cli, args); s == ExitCodeOk {
			t.Errorf("it is Expected that it would be failed. Exit code is %v. Too much arguments", s)
		}
	})
}
