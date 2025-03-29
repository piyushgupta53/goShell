package builtins

import (
	"fmt"
	"os"
)

func Type(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "type: missing operand")
		return
	}

	for _, name := range args {
		if IsBuiltIn(name) {
			fmt.Printf("%s is a shell builtin\n", name)
		} else {
			// Will handle executables in Step 10
			fmt.Fprintf(os.Stderr, "%s: not found\n", name)
		}
	}

}
