package builtins

import (
	"fmt"
	"os"
)

// Cd changes the working directory
func Cd(args []string) {
	var targetDir string

	if len(args) == 0 || args[0] == "~" {
		// No args or ~ => go to home dir
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cd: cannot change home directory")
			return
		}

		targetDir = home
	} else {
		// We'll support args in the next steps
		fmt.Fprintln(os.Stderr, "cd: unsupported usage — only 'cd' and 'cd ~' work for now")
		return
	}

	if err := os.Chdir(targetDir); err != nil {
		fmt.Fprintf(os.Stderr, "cd: %v\n", err)
	}
}
