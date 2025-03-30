package parser

import (
	"strings"
)

// Parse splits the input string into a Command (command + args)
// e.g. input: "ls -ls/tmp" => Command{Name: "ls", Args: ["-la", "/tmp"]}
func Parse(input string) *Command {
	tokens := tokenize(input)

	if len(tokens) == 0 {
		return nil
	}

	return &Command{
		Name: tokens[0],
		Args: tokens[1:],
	}
}

func tokenize(input string) []string {
	var tokens []string
	var current strings.Builder
	inSingleQuote := false
	inDoubleQuote := false
	escaped := false

	for i := 0; i < len(input); i++ {
		ch := input[i]

		switch {
		case inSingleQuote:
			if ch == '\'' {
				// End of single quote
				inSingleQuote = false
			} else {
				current.WriteByte(ch)
			}

		case inDoubleQuote:
			if escaped {
				switch ch {
				case '"', '\\':
					current.WriteByte(ch)
				default:
					current.WriteByte('\\')
					current.WriteByte(ch)
				}
				escaped = false
			} else if ch == '\\' {
				escaped = true
			} else if ch == '"' {
				inDoubleQuote = false
			} else {
				current.WriteByte(ch)
			}

		default:
			if escaped {
				current.WriteByte(ch)
				escaped = false
				continue
			}

			switch ch {
			case '\\':
				escaped = true
			case '\'':
				inSingleQuote = true
			case '"':
				inDoubleQuote = true
			case ' ', '\t':
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
			default:
				current.WriteByte(ch)
			}
		}

	}

	// flush last tokens
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}
