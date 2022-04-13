package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args

	for i := len(params) - 1; i > 0; i-- {
		str := []rune(params[i])
		for j := 0; j < len(str); j++ {
			z01.PrintRune(str[j])
		}
		z01.PrintRune('\n')
	}
}
