# GoShell

A lightweight command-line shell implementation written in Go. This project provides a simple yet functional shell environment that supports both built-in commands and external program execution.

## Features

### Built-in Commands

- `cd` - Change directory
- `echo` - Print arguments to standard output
- `exit` - Exit the shell
- `pwd` - Print working directory
- `type` - Display information about command type

### Command Line Features

- **Command Completion**: Press TAB to get suggestions for:
  - Built-in commands
  - Executable files in PATH
  - File and directory paths
- **History**: Access command history using up/down arrow keys
- **Line Editing**: Basic line editing capabilities
- **Prompt Customization**: Configurable shell prompt

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
├── cmd/
│   └── shell/       # Main entry point
├── internal/
│   ├── builtins/    # Built-in command implementations
│   ├── completion/  # Command completion functionality
│   ├── executor/    # Command execution logic
│   ├── parser/      # Command parsing and tokenization
│   ├── repl/        # Read-Eval-Print Loop implementation
│   └── utils/       # Utility functions
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

# Command completion
# Press TAB to see suggestions for:
ls <TAB>           # Lists files and directories
cd <TAB>           # Lists directories
python <TAB>       # Lists Python files
```

## Features in Detail

### Command Completion

- Press TAB to get suggestions for commands and files
- Supports completion for:
  - Built-in shell commands
  - System executables in PATH
  - File and directory paths
- Shows multiple suggestions when multiple matches are found

### Error Handling

The shell provides clear error messages for common scenarios:

- Command not found
- Execution failures
- Invalid arguments
- Permission issues
- File/directory not found

## Contributing

Feel free to submit issues and enhancement requests!
