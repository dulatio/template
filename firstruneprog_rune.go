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
	if ln == 1 {
		var r rune
		for _, c := range arg[0] {
			r = c
			break
		}
		z01.PrintRune(r)
		z01.PrintRune('\n')
	}
}
