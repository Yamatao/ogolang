package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env
func RunCmd(args []string, env Environment) (retCode int, err error) {
	var cmd *exec.Cmd
	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else if len(args) == 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		return 0, errors.New("args is empty")
	}
	if cmd == nil {
		return 0, errors.New("os.exec.Command returned nil")
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Env = env.AsArray()

	err = cmd.Start()
	if err != nil {
		return 0, fmt.Errorf("Failed to run a process: %v", err)
	}

	state, err := cmd.Process.Wait()
	if err != nil {
		return 0, fmt.Errorf("Failed to wait process exit: %v", err)
	}

	return state.ExitCode(), nil
}
