package executor

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/piyushgupta53/goShell/internal/builtins"
	"github.com/piyushgupta53/goShell/internal/parser"
)

// Execute runs a parsed command as an external program
func Execute(cmd *parser.Command) {
	if cmd == nil || cmd.Name == "" {
		return
	}

	// Builtin check
	if builtins.IsBuiltIn(cmd.Name) {
		builtins.RunBuiltIn(cmd.Name, cmd.Args)
		return
	}

	// Check if command exists in PATH
	path, err := exec.LookPath(cmd.Name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "command not found: %s\n", cmd.Name)
		return
	}

	// Setup command with args
	execCmd := exec.Command(path, cmd.Args...)
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Execute
	if err := execCmd.Run(); err != nil {
		// Note: this includes exit code errors, permissions etc,
		fmt.Fprintf(os.Stderr, "%s: execution failed: %v\n", cmd.Name, err)
	}
}
