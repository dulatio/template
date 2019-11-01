package main

import "github.com/01-edu/z01"

func Rot14(str string) string {
	var result string
	arrayRune := []rune(str)
	for index, x := range arrayRune {
		if (x >= 'A' && x < 'M') || (x >= 'a' && x < 'm') {
			arrayRune[index] = x + 14
		} else if (x >= 'M' && x <= 'Z') || (x >= 'm' && x <= 'z') {
			arrayRune[index] = x - 12
		}

		result += string(arrayRune[index])

	}
	return result
}

func main() {
	result := Rot14("Hello How are You")
	arrayRune := []rune(result)

	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
