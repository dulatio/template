package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		for _, c := range arg[0] {
			if c >= 'a' && c <= 'z' {
				c += 13
				if c > 'z' {
					c = rune(c - 26)
				}
			}
			if c >= 'A' && c <= 'Z' {
				c += 13
				if c > 'Z' {
					c = rune(c - 26)
				}
			}
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
