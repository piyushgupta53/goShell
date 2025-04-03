package parser

import (
	"strings"
)

// Parse splits the input string into a Command (command + args)
// e.g. input: "ls -ls/tmp" => Command{Name: "ls", Args: ["-la", "/tmp"]}
func Parse(input string) *Command {
	rawTokens := tokenize(input)
	if len(rawTokens) == 0 {
		return nil
	}

	var name string
	var args []string
	var redirs []Redirection

	i := 0
	for i < len(rawTokens) {
		token := rawTokens[i]

		// Handle redirections
		if token == ">" || token == ">>" || token == "2>" || token == "2>>" {
			if i+1 >= len(rawTokens) {
				//missing filename
				return nil
			}

			target := rawTokens[i+1]
			shouldAppend := token == ">>" || token == "2>>"
			fd := 1
			if strings.HasPrefix(token, "2") {
				fd = 2
			}
			redirs = append(redirs, Redirection{
				Fd:     fd,
				Append: shouldAppend,
				Target: target,
			})
			i += 2 // skip operator and filename
			continue
		}

		// Handle command name and arguments
		if name == "" {
			name = token
		} else {
			args = append(args, token)
		}
		i++
	}

	return &Command{
		Name:         name,
		Args:         args,
		Redirections: redirs,
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
				// Always treat next character literally
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

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}
