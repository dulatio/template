package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	tmp := ""
	l := 0
	if len(arg) == 1 {
		for _, c := range arg[0] {
			if c == ' ' && tmp != "" {
				tmp = ""
				l++
			} else if c != ' ' {
				tmp = tmp + string(c)
			}
		}
		if tmp != "" {
			l++
			tmp = ""
		}
		ans := make([]string, l)
		l = 0
		for _, c := range arg[0] {
			if c == ' ' && tmp != "" {
				ans[l] = tmp
				l++
				tmp = ""
			} else if c != ' ' {
				tmp = tmp + string(c)
			}
		}
		if tmp != "" {
			ans[l] = tmp
			l++
			tmp = ""
		}
		ok := true
		for _, c := range ans {
			if ok {
				ok = false
				for _, w := range c {
					z01.PrintRune(w)
				}
			} else {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				tmp = ""
			}
			ans := make([]string, l)
			l = 0
			for _, c = range arg[0] {
				if c == ' ' && tmp != "" {
					ans[l] = tmp
					l++
					tmp = ""
				} else if c != ' ' {
					tmp = tmp + string(c)
				}
			}
			if tmp != "" {
				ans[l] = tmp
				l++
				tmp = ""
			}
			ok := true
			for _, c := range ans {
				if ok {
					ok = false
					for _, w := range c {
						z01.PrintRune(w)
					}
				} else {
					z01.PrintRune(' ')
					z01.PrintRune(' ')
					z01.PrintRune(' ')

					for _, w := range c {
						z01.PrintRune(w)
					}
				}
			}
		}
		z01.PrintRune('\n')
	}

