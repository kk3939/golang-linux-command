package main

import (
	"fmt"
	"os"
)

func hasNoArgs(args []string) error {
	if len(os.Args) == 1 {
		return fmt.Errorf("error: %s", "Specify linux command to a first argument")
	}
	return nil
}

func main() {
	args := os.Args
	if err := hasNoArgs(args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch args[1] {
	case "ls":
		fmt.Println("test")
	default:
		fmt.Println("error: Command is not implemented.")
		os.Exit(1)
	}
}
