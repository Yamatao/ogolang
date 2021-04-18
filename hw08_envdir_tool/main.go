package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("envdir <dir-path> <command> [args...] - run a command with a specific environment read from the given directory.")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Wrong number of arguments. Use --help.\n")
		return
	}
	if os.Args[0] == "--help" {
		printUsage()
		return
	}

	env, err := ReadDir(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read environment directory: %v\n", err)
		return
	}

	retCode, err := RunCmd(os.Args[2:], env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run a command: %v\n", err)
		return
	}

	os.Exit(retCode)
}
