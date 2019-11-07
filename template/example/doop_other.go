package main

import (
	"fmt"
	"os"
)

func chartoint(c rune) int {
	return int(c) - 48
}

func Atoi(s string) int {
	str := []rune(s)
	answer := 0
	sign := false
	positive := true

	length := 0

	for range s {
		length++
	}

	for i := range str {
		if answer == 0 && str[i] == '0' {
			answer = 0
		} else if answer == 0 && str[i] == '+' {
			if sign == true {
				return 0
			}
			sign = true
		} else if answer == 0 && str[i] == '-' {
			if sign == true {
				return 0
			}
			sign = true
			positive = false
		} else if str[i] < '0' || str[i] > '9' {
			return 0
		} else {
			ten := 1
			for j := i; j < length-1; j++ {
				ten = ten * 10
			}
			answer = answer + chartoint(str[i])*ten
		}
	}
	if positive == true {
		return answer
	} else {
		return -answer
	}
}

func modulo(a int, b int) {
	if b == 0 {
		fmt.Println("No Modulo by 0")
		return
	}
	fmt.Println(a % b)
}

func division(a int, b int) {
	if b == 0 {
		fmt.Println("No division by 0")
		return
	}
	fmt.Println(a / b)
}

func multiplication(a int, b int) {
	fmt.Println(a * b)
}

func subtraction(a int, b int) {
	fmt.Println(a - b)
}

func addition(a int, b int) {
	fmt.Println(a + b)
}

func callFunc(a, b, c string) {
	for _, char := range a {
		if (char < 48 || char > 57) && rune(char) != '-' {
			fmt.Println("0")
			return
		}
	}
	for _, char := range c {
		if (char < 48 || char > 57) && rune(char) != '-' {
			fmt.Println("0")
			return
		}
	}

	nbr1 := Atoi(a)
	nbr2 := Atoi(c)

	if b == "+" {
		addition(nbr1, nbr2)
	} else if b == "-" {
		subtraction(nbr1, nbr2)
	} else if b == "*" {
		multiplication(nbr1, nbr2)
	} else if b == "/" {
		division(nbr1, nbr2)
	} else if b == "%" {
		modulo(nbr1, nbr2)
	} else {
		fmt.Println("0")
		return
	}
}

func main() {
	nbrOfArguments := -1
	for range os.Args {
		nbrOfArguments++
	}

	if nbrOfArguments != 3 {
		return
	}
	callFunc(os.Args[1], os.Args[2], os.Args[3])
}
