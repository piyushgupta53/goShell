package builtins

// IsBuiltIn checks if a command is a builtin
func IsBuiltIn(name string) bool {
	switch name {
	case "exit", "echo":
		return true
	default:
		return false
	}
}

// RunBuiltIn dispatches the builtin function
func RunBuiltIn(name string, args []string) {
	switch name {
	case "exit":
		Exit(args)
	case "echo":
		Echo(args)
	}
}
