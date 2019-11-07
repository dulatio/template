package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	str := []rune(s)
	length := 0
	for range str {
		length++
	}
	return str[length-1]

}

func main() {
	z01.PrintRune(LastRune("Hello"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune('\n')
}
