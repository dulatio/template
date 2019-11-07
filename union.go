package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 2 {
		st := map[rune]bool{}
		ans := arg[0] + arg[1]
		cur := ""
		for _, c := range ans {
			if !st[c] {
				cur += string(c)
				st[c] = true
			}
		}
		fmt.Print(cur)
	}
	fmt.Println()
}
