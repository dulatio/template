package main

import (
	"os"

	"github.com/01-edu/z01"
)

func err(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return true
		}
	}
	return false
}
func BasicAtoi2(s string) int {
	x := 0
	for _, c := range s {
		k := 0
		for i := '1'; i <= c; i++ {
			k++
		}
		x = x*10 + k
	}
	return x
}
func main() {
	args := os.Args[1:]
	k := 0
	for range args {
		k++
	}
	ans := "str"
	if k == 1 {
		if err(args[0]) {
			for _, c := range ans {
				z01.PrintRune(c)
			}
		} else {
			n := BasicAtoi2(args[0])
			pw := 1
			ok := false
			for pw <= n {
				if pw == n {
					ok = true
				}
				pw *= 2
			}
			if ok {
				ans = "true"
			} else {
				ans = "false"
			}
			for _, c := range ans {
				z01.PrintRune(c)
			}
		}
	}
	z01.PrintRune('\n')
}
