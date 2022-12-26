package main

import (
	"fmt"
	"os"
	"strings"
)

func hasNoArgs() error {
	if len(os.Args) == 1 {
		return fmt.Errorf("golang-linux-command: %s", "specify linux command to a first argument!")
	}
	return nil
}

func main() {
	if err := hasNoArgs(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ls":
		p, err := os.Getwd()
		if err != nil {
			fmt.Println(fmt.Errorf("golang-linux-command: %w", err))
			os.Exit(1)
		}

		files, err := os.ReadDir(p)
		if err != nil {
			fmt.Println(fmt.Errorf("golang-linux-command: %w", err))
			os.Exit(1)
		}

		for _, file := range files {
			// skip hidden file
			if strings.HasPrefix(file.Name(), ".") {
				continue
			}
			fmt.Println(file.Name())
		}

	default:
		fmt.Println(fmt.Errorf("golang-linux-command: %s", "command is not implemented!"))
		os.Exit(1)
	}
}
