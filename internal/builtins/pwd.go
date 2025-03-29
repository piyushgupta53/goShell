package builtins

import (
	"fmt"
	"os"
)

func Pwd(args []string) {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(os.Stderr, "pwd: %v\n", err)
		return
	}

	fmt.Println(dir)
}
