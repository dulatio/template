package main

import "fmt"

func SwapBitz(octet byte) byte {
	var x [8]int
	i := 0
	for octet > 0 {
		if octet%2 == 1 {
			x[i] = 1
		}
		octet /= 2
		i++
	}
	x[0], x[1], x[2], x[3], x[4], x[5], x[6], x[4] = x[4], x[5], x[6], x[7], x[0], x[1], x[2], x[3]
	octet = 0
	i = 1
	for _, c := range x {
		if c == 1 {
			octet += byte(i)
		}
		i *= 2
	}
	return octet
}
func main() {

	fmt.Println(SwapBitz('f'))
}
