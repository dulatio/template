package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	ln := 0
	for c := range arg {
		ln = c + 1
	}
	if ln == 2 {
		s := arg[0]
		ln_s := 0
		for i := range s {
			ln_s = i + 1
		}
		t := arg[1]
		cnt := 0
		for _, c := range t {
			if c == rune(s[cnt]) {
				cnt++
			}
			if cnt == ln_s {
				for _, w := range s {
					z01.PrintRune(w)
				}
				// fmt.Print(s)
				break
			}
		}
	}
	z01.PrintRune('\n')
	// fmt.Println()
}
