package builtins

import (
	"fmt"
	"strings"
)

func Echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}
