package completion

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/piyushgupta53/goShell/internal/builtins"
)

var cachedExecutables = []string{}

func init() {
	cachedExecutables = loadExecutablesFromPATH()
}

func Trigger(line string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}
	current := parts[len(parts)-1]

	var matches []string

	// First token: command
	if len(parts) == 1 {
		for _, b := range builtins.AllBuiltIns() {
			if strings.HasPrefix(b, current) {
				matches = append(matches, b)
			}
		}
		for _, cmd := range cachedExecutables {
			if strings.HasPrefix(cmd, current) {
				matches = append(matches, cmd)
			}
		}
	} else {
		// Argument: complete file path
		matches = completeFilePaths(current)
	}

	switch len(matches) {
	case 0:
		// nothing
	case 1:
		suffix := matches[0][len(current):]
		fmt.Print(suffix)
	default:
		// multiple: format and show suggestions
		fmt.Println()
		fmt.Println(strings.Join(matches, "    "))
		// reprint prompt and current line
		fmt.Print("> ", line)
	}
}

func loadExecutablesFromPATH() []string {
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	seen := make(map[string]bool)
	var execs []string

	for _, dir := range paths {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			name := entry.Name()
			full := filepath.Join(dir, name)
			if entry.Type().IsRegular() || entry.Type()&os.ModeSymlink != 0 {
				if seen[name] {
					continue
				}
				if isExecutable(full) {
					execs = append(execs, name)
					seen[name] = true
				}
			}
		}
	}

	return execs
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0111 != 0
}

func completeFilePaths(prefix string) []string {
	dir := "."
	base := prefix

	if strings.Contains(prefix, "/") {
		dir = filepath.Dir(prefix)
		base = filepath.Base(prefix)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var results []string
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, base) {
			full := filepath.Join(dir, name)
			if entry.IsDir() {
				full += "/"
			}
			results = append(results, full)
		}
	}

	return results
}
