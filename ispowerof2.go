package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	ln := 0
	for range arg {
		ln++
	}
	if ln == 1 {
		ans := "true"
		x, err := strconv.Atoi(arg[0])
		if err != nil {
			ans = err.Error()
		} else {
			if x == 0 {
				ans = "false"
			}
			for x > 1 {
				if x%2 == 1 {
					ans = "false"
					break
				}
				x /= 2
			}
		}
		for _, c := range ans {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
