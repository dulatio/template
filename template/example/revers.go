package main

import (
 "os"

 "github.com/01-edu/z01"
)

func check(s string) bool {
 for _, c := range s {
  if c != '0' && c != '1' {
   return false
  }
 }
 return true
}

func main() {
 arg := os.Args[1:]
 ln := 0
 for range arg {
  ln++
 }
 if ln == 1 {
  if check(arg[0]) {
   ln = 0
   for i := range arg[0] {
    ln = i
   }
   for i := ln; i >= 0; i-- {
    z01.PrintRune(rune(arg[0][i]))
   }
  } else {
   ans := "The" + arg[0] + " is argument"
   for _, c := range ans {
    z01.PrintRune(c)
   }
  }
 }
 z01.PrintRune('\n')
}