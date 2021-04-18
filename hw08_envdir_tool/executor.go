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
	switch {
	case len(args) > 1:
		cmd = exec.Command(args[0], args[1:]...) //nolint:gosec
	case len(args) == 1:
		cmd = exec.Command(args[0], args[1:]...) //nolint:gosec
	default:
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
		return 0, fmt.Errorf("failed to run a process: %w", err)
	}

	state, err := cmd.Process.Wait()
	if err != nil {
		return 0, fmt.Errorf("failed to wait process exit: %w", err)
	}

	return state.ExitCode(), nil
}
