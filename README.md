# GoShell

A lightweight command-line shell implementation written in Go. This project provides a simple yet functional shell environment that supports both built-in commands and external program execution.

## Features

### Built-in Commands

- `cd` - Change directory
- `echo` - Print arguments to standard output
- `exit` - Exit the shell
- `pwd` - Print working directory
- `type` - Display information about command type

### External Command Execution

- Execute any command available in the system PATH
- Support for absolute and relative paths
- Proper handling of command arguments
- Standard I/O redirection (stdin, stdout, stderr)

### Command Parsing

- Robust command parsing with support for:
  - Single quotes
  - Double quotes
  - Escape characters
  - Whitespace handling
  - Multiple arguments

## Project Structure

```
goShell/
├── internal/
│   ├── builtins/    # Built-in command implementations
│   ├── executor/    # Command execution logic
│   └── parser/      # Command parsing and tokenization
└── main.go          # Main entry point
```

## Building and Running

To build and run the shell:

```bash
go build
./go-shell
```

## Usage

The shell supports both built-in commands and external programs. Here are some examples:

```bash
# Built-in commands
pwd
cd /path/to/directory
echo "Hello, World!"
exit

# External commands
ls -la
cat file.txt
python script.py
```

## Error Handling

The shell provides clear error messages for common scenarios:

- Command not found
- Execution failures
- Invalid arguments
- Permission issues
