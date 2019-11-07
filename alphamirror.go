package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		for _, i := range arg[0] {
			c := i
			if c >= 'a' && c <= 'z' {
				c = rune(219 - c)

			} else if c >= 'A' && c <= 'Z' {
				c = rune(155 - c)

			}
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
