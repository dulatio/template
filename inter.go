package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 2 {
		st := map[rune]bool{}
		st2 := map[rune]bool{}
		for _, c := range arg[1] {
			st[c] = true
		}
		for _, c := range arg[0] {
			if !st2[c] && st[c] {
				z01.PrintRune(c)
				st2[c] = true
			}
		}
	}
	z01.PrintRune('\n')
}
