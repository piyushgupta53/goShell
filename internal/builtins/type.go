package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func Type(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "type: missing operand")
		return
	}

	for _, name := range args {
		if IsBuiltIn(name) {
			fmt.Printf("%s is a shell builtin\n", name)
			continue
		}

		// Check if executable is in PATH
		if path, err := exec.LookPath(name); err == nil {
			fmt.Printf("%s is %s\n", name, path)
		} else {
			fmt.Fprintf(os.Stderr, "%s: not found\n", name)
		}
	}

}
