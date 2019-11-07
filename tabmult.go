package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		x := 0
		for _, c := range arg[0] {
			x = x*10 + int(c-'0')
		}
		cnt := 1
		for i := '1'; i <= '9'; i++ {
			cur := x * cnt
			ans := string(i) + " * " + arg[0] + " = "
			add := ""
			y := cur
			if y == 0 {
				add = "0"
			}
			for y > 0 {
				add = string(rune(y%10+'0')) + add
				y /= 10
			}
			ans = ans + add
			for _, c := range ans {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			cnt++
		}
	}
}
