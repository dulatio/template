package main

import "fmt"

func StrLen(str string) int {
	length := 0
	strc := []rune(str)
	for index := range strc {
		length = index + 1
	}
	return length
}

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
