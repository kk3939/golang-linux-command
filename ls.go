package main

import (
	"fmt"
	"os"
	"strings"
)

func ls() {
	var path string
	if len(os.Args) == 3 {
		// When path argument is given
		path = os.Args[2]
	} else if len(os.Args) == 2 {
		// When no argument
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(fmt.Errorf("golang-linux-command: %w", err))
			os.Exit(1)
		}
		path = currentDir
	} else {
		// usage that is not considered
		fmt.Println(fmt.Errorf("golang-linux-command: %s", "expected arguments are too much"))
		os.Exit(1)
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
