package environments

import (
	"os"
	"os/exec"
)

// getShell : returns default system shell, if $SHELL is not set returns "/bin/bash"
func getShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}

	return shell
}

// SpawnShell : Spawns default system shell with injected variables
func SpawnShell() {

	shell := getShell()

	cmd := exec.Command(shell)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()
	cmd.Wait()
}
