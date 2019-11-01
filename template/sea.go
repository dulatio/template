    package main

    import (
    	"os"

    	"github.com/01-edu/z01"
    )

    func main() {
    	args := os.Args[1:]
    	if len(args) == 3 {
    		string0 := args[0]
    		string1 := args[1]
    		string2 := args[2]

    		str_result := ""
    		for i := 0; i < len(string0); i++ {
    			if byte(string0[i]) == byte(string1[0]) {
    				str_result = str_result + string(string2[0])
    			} else {
    				str_result = str_result + string(string0[i])
    			}

    		}

    		for _, value := range str_result {
    			z01.PrintRune(value)
    		}

    	}
    	z01.PrintRune(10)
    }