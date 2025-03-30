package builtins

import (
	"fmt"
	"os"
	"strings"
)

// Cd changes the working directory
func Cd(args []string) {
	var targetDir string

	switch {
	case len(args) == 0 || args[0] == "~":
		// No args or ~ => go to home dir
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cd: cannot change home directory")
			return
		}

		targetDir = home

	case strings.HasPrefix(args[0], "/"):
		// Absolute path
		targetDir = args[0]
	default:
		// Relative path
		targetDir = args[0]
	}

	// Try to change directory
	if err := os.Chdir(targetDir); err != nil {
		fmt.Fprintf(os.Stderr, "cd: %v\n", err)
	}
}
