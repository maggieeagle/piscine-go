package main

import (
	"github.com/01-edu/z01"
)

func main() {
	digits := []rune("0123456789\n")
	for i := 0; i < 11; i++ {
		z01.PrintRune(digits[i])
	}
}
