package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	ln := 0
	for range arg {
		ln++
	}
	if ln == 2 {
		if arg[0] != arg[1] {
			if arg[0] < arg[1] {
				z01.PrintRune('-')
			}
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
	}
	z01.PrintRune('\n')
}
