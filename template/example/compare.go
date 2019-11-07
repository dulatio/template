package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	args := os.Args

	for range args {
		count++
	}
	if count != 3 {
		z01.PrintRune(10)
		return
	}

	args1 := os.Args[1]
	args2 := os.Args[2]
	for range args1 {

		if args1 < args2 {
			z01.PrintRune('-')
			z01.PrintRune('1')
			z01.PrintRune(10)
		}
		if args1 > args2 {
			z01.PrintRune('1')
			z01.PrintRune(10)
		}
		if args1 == args2 {
			z01.PrintRune('0')
			z01.PrintRune(10)
		}
		for range os.Args[2] {

		}
		return
	}
}
