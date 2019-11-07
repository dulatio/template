package main

import "fmt"

func Max(arr []int) int {
	maxV := 0
	for _, v := range arr {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	arrInt := []int{23, 123, 1, 11, 55, 93}
	max := Max(arrInt)
	fmt.Println(max)
}
