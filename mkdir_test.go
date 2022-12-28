package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"testing"
)

func init_mkdir_test() func() {
	test := func(name string, perm fs.FileMode) error {
		b := new(bytes.Buffer)
		if _, err := fmt.Fprintf(b, "Testing: %s\n", "Success!!"); err != nil {
			return err
		}
		return nil
	}
	resetF := SetCreateDir(test)
	return resetF
}

func Test_mkdir(t *testing.T) {
	t.Run("one_argument", func(t *testing.T) {
		resetF := init_mkdir_test()
		defer resetF()
		cli, args, _, _ := setup_test("golang-linux-command mkdir temp")
		if s := mkdir(cli, args); s == ExitCodeError {
			t.Errorf("Exit code is %v", s)
		}
	})

	t.Run("no_argument", func(t *testing.T) {
		resetF := init_mkdir_test()
		defer resetF()
		cli, args, _, _ := setup_test("golang-linux-command mkdir")
		if s := mkdir(cli, args); s == ExitCodeOk {
			t.Errorf("Exit code is %v", s)
		}
	})

	t.Run("one_more_argument", func(t *testing.T) {
		resetF := init_mkdir_test()
		defer resetF()
		cli, args, _, _ := setup_test("golang-linux-command mkdir test test")
		if s := ls(cli, args); s == ExitCodeOk {
			t.Errorf("it is Expected that it would be failed. Exit code is %v. Too much arguments", s)
		}
	})
}
