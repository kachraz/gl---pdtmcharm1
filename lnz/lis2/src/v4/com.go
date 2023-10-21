/*
This is for implmenting tea.ExecCmd in v.go
*/

package src

import (
	"io"
	"os"
	"os/exec"
)

type MyCommand struct {
	command  string
	selected string
}

func (c MyCommand) Run() error {
	cmd := exec.Command(c.command, c.selected)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func (c MyCommand) SetStdin(r io.Reader) {
	// Set the standard input for the command, if needed
}

func (c MyCommand) SetStdout(w io.Writer) {
	// Set the standard output for the command, if needed
}

func (c MyCommand) SetStderr(w io.Writer) {
	// Set the standard error for the command, if needed
}
