package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	args := os.Args[1:]

	for range args {
		count++
	}
	if count > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if count == 0 {
		z01.PrintRune(10)
	} else {
		a := []rune(args[0])
		for i := 0; i < count; i++ {
			n := true
			for i, c := range a {
				if c >= 'A' && c <= 'Z' && n == true {
					n = false
				} else if c >= 'A' && c <= 'Z' && n == false {
					a[i] = c + 32
					n = false
				} else if a[i] >= '0' && a[i] <= '9' {
					n = false
				} else if c >= 'a' && c <= 'z' && n == false {
					n = false
				} else if c >= 'a' && c <= 'z' && n == true {
					a[i] = c - 32
					n = false
				} else {
					n = true
				}
			}
		}
		for _, value := range a {
			z01.PrintRune(value)
		}
		z01.PrintRune(10)
	}
}
