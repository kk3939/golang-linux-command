package main

import (
	"fmt"
	"os"
)

func cat() {
	if len(os.Args) > 3 {
		fmt.Println(fmt.Errorf("golang-linux-command: %s", "expected arguments are too much"))
		os.Exit(1)
	}

	// When no argument
	if len(os.Args) == 2 {
		fmt.Println(fmt.Errorf("golang-linux-command: %s", "argument is required!"))
		os.Exit(1)
	}

	file, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(fmt.Errorf("golang-linux-command: %w", err))
		os.Exit(1)
	}
	fmt.Println(string(file))
}
