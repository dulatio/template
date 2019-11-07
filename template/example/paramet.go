    package main

    import (
    	"os"

    	"github.com/01-edu/z01"
    )

    func main() {
    	args := os.Args[1:]
    	count := len(args)
    	ans := ""
    	if count > 0 {
    		for count >= 1 {
    			ans = string(rune(count%10+'0')) + ans
    			count /= 10
    		}
    	} else if count == 0 {
    		z01.PrintRune('0')
    	}
    	for _, c := range ans {
    		z01.PrintRune(c)
    	}
    	z01.PrintRune(10)
    }