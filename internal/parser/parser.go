package parser

import "strings"

// Parse splits the input string into a Command (command + args)
// e.g. input: "ls -ls/tmp" => Command{Name: "ls", Args: ["-la", "/tmp"]}
func Parse(input string) *Command {
	tokens := strings.Fields(input)

	if len(tokens) == 0 {
		return nil
	}

	return &Command{
		Name: tokens[0],
		Args: tokens[1:],
	}
}
