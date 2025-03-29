package builtins

import (
	"fmt"
	"os"
	"strconv"
)

// Exit handles the 'exit' builtin command
func Exit(args []string) {
	code := 0

	if len(args) > 0 {
		// try to parse the optional exit code
		if n, err := strconv.Atoi(args[0]); err != nil {
			code = n
		} else {
			fmt.Fprintf(os.Stderr, "exit: invalid exit code: %s\n", args[0])
		}
	}

	os.Exit(code)
}
