package main

import (
	"fmt"
	"os"
	"strings"
)

func ls() {
	if len(os.Args) > 3 {
		fmt.Println(fmt.Errorf("golang-linux-command: %s", "expected arguments are too much"))
		os.Exit(1)
	}

	var path string
	// When path argument is given
	if len(os.Args) == 3 {
		path = os.Args[2]
	}
	// When no argument
	if len(os.Args) == 2 {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(fmt.Errorf("golang-linux-command: %w", err))
			os.Exit(1)
		}
		path = currentDir
	}

	files, err := os.ReadDir(path)
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
}
