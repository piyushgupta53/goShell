package completion

import (
	"fmt"
	"strings"

	"github.com/piyushgupta53/goShell/internal/builtins"
)

func Trigger(line string) {
	// extract the last word before the cursor(naive split)
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}

	current := parts[len(parts)-1]

	matches := []string{}
	for _, b := range builtins.AllBuiltIns() {
		if strings.HasPrefix(b, current) {
			matches = append(matches, b)
		}
	}

	switch len(matches) {
	case 0:
		// no match - do nothing
	case 1:
		// single match - suggestion completion
		suffix := matches[0][len(current):]
		fmt.Print(suffix)
	case 2, 3, 4, 5:
		// few matches - show suggestions
		fmt.Print("\r")             // Return to start of line
		fmt.Print("\033[K")         // Clear the line
		fmt.Print("> ", line, "\n") // Print prompt and current line
		for _, match := range matches {
			fmt.Println(match)
		}
		fmt.Print("> ", line) // Reprint prompt and current line
	default:
		fmt.Print("\r")             // Return to start of line
		fmt.Print("\033[K")         // Clear the line
		fmt.Print("> ", line, "\n") // Print prompt and current line
		fmt.Print("[too many matches]\n")
		fmt.Print("> ", line) // Reprint prompt and current line
	}
}
