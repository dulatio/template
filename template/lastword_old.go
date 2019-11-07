package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	length := 0
	for range os.Args {
		length++
	}
	if length != 2 {
		return
	}

	count := 0
	pos := -1
	slice := []rune(os.Args[length-1])
	for range slice {
		count++
	}
	for index, char := range slice {
		if char == ' ' && index != count-1 {
			if slice[index+1] != ' ' {
				pos = index
			}
		}
	}
	start := 0
	if pos != -1 {
		start = pos + 1
	}
	for i := start; i < count; i++ {
		z01.PrintRune(slice[i])
	}
	z01.PrintRune('\n')

}
