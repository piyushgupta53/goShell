package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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

	var execPath string
	var err error

	// If command has a slash, treat it as a path - don't use LookPath
	if strings.ContainsRune(cmd.Name, '/') {
		execPath = cmd.Name
	} else {
		// Search path
		execPath, err = exec.LookPath(cmd.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "command not found: %s\n", cmd.Name)
			return
		}
	}

	// Check if command exists in PATH

	// Setup command with args
	execCmd := exec.Command(execPath, cmd.Args...)
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Execute
	if err := execCmd.Run(); err != nil {
		// Note: this includes exit code errors, permissions etc,
		fmt.Fprintf(os.Stderr, "%s: execution failed: %v\n", cmd.Name, err)
	}
}
