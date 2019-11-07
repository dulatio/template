package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		tmp := ""
		for i := len(arg[0]) - 1; i >= 0; i-- {
			if arg[0][i] == ' ' && tmp != "" {
				break
			} else if arg[0][i] != ' ' {
				tmp = string(arg[0][i]) + tmp
			}
		}
		for _, c := range tmp {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')

}
