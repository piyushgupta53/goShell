package repl

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/piyushgupta53/goShell/internal/completion"
	"github.com/piyushgupta53/goShell/internal/executor"
	"github.com/piyushgupta53/goShell/internal/parser"
	"golang.org/x/term"
)

func Start() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error setting raw mode: %v\n", err)
	}

	defer term.Restore(int(syscall.Stdin), oldState)

	fmt.Print("> ")

	line := make([]byte, 0, 1024)
	buf := make([]byte, 1)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}

		b := buf[0]

		switch b {
		case '\r', '\n':
			// enter key
			fmt.Print("\n")
			input := string(line)
			line = line[:0] // reset buffer

			if strings.TrimSpace(input) == "" {
				fmt.Print("> ")
				continue
			}

			cmd := parser.Parse(input)

			if cmd != nil {
				executor.Execute(cmd)
			}
			fmt.Print("> ")

		case 127:
			// backspace
			if len(line) > 0 {
				line = line[:len(line)-1]
				fmt.Print("\b \b")
			}

		case '\t':
			// tab pressed
			completion.Trigger(string(line))
			fmt.Print("> ", string(line)) // reprint prompt and line

		default:
			// normal character
			line = append(line, b)
			fmt.Printf("%c", b)
		}
	}
}
