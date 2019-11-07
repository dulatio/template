package main

import "github.com/01-edu/z01"

func Rot13(str string) string {
	var result string
	arrayRune := []rune(str)

	for index, x := range arrayRune {
		if (x >= 'a' && x < 'm') || (x >= 'A' && x < 'M') {
			arrayRune[index] = x + 13
		} else if (x >= 'm' && x < 'z') || (x >= 'M' && x < 'Z') {
			arrayRune[index] = x - 13
		}
		result += string(arrayRune[index])

	}
	return result
}

func main() {
	result := Rot13("hello there")
	arrayRune := []rune(result)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
