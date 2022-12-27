package main

import (
	"fmt"
	"os"
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
		ls()
	case "cat":
		cat()
	default:
		fmt.Println(fmt.Errorf("golang-linux-command: %s", "command is not implemented!"))
		os.Exit(1)
	}
}
