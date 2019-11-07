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
		ln = 0
		for range arg[0] {
			ln++
		}
		ans := ""
		if ln == 0 {
			ans = "0"
		}
		for ln > 0 {
			ans = string(rune(ln%10+'0')) + ans
			ln /= 10
		}
		for _, c := range ans {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
