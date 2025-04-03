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

	// Setup command with args
	execCmd := exec.Command(execPath, cmd.Args...)
	execCmd.Stdin = os.Stdin

	// Don't set stdout/stderr by default - wait to see if we have redirections
	defaultStdout := os.Stdout
	defaultStderr := os.Stderr

	// Open files for redirection
	var redirFiles []*os.File
	for _, r := range cmd.Redirections {
		// Set appropriate flags based on append mode
		var flag int
		if r.Append {
			flag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
		} else {
			flag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
		}

		f, err := os.OpenFile(r.Target, flag, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "redirection error: %v\n", err)
			return
		}

		redirFiles = append(redirFiles, f)

		switch r.Fd {
		case 1:
			defaultStdout = f
		case 2:
			defaultStderr = f
		}
	}

	// Set the final output destinations
	execCmd.Stdout = defaultStdout
	execCmd.Stderr = defaultStderr

	// Run the command
	err = execCmd.Run()

	// Close redirection files
	for _, f := range redirFiles {
		if err := f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "error closing redirection: %v\n", err)
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: execution failed: %v\n", cmd.Name, err)
	}
}
