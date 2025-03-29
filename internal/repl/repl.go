package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/piyushgupta53/goShell/internal/executor"
	"github.com/piyushgupta53/goShell/internal/parser"
)

func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			// handle EOF
			if err == io.EOF {
				fmt.Println("\nExiting shell...")
				os.Exit(0)
			}

			// other errors
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		}

		// trim input
		line = strings.TrimSpace(line)

		// empty line? loop again
		if line == "" {
			continue
		}

		cmd := parser.Parse(line)
		if cmd == nil {
			continue
		}

		executor.Execute(cmd)
	}
}
