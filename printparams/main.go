package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	str := []rune(params[0])
	var index int = 0
	for i := 0; i < len(str); i++ {
		if str[i] == '/' {
			index = i
		}
	}
	for i := index + 1; i < len(str); i++ {
		z01.PrintRune(str[i])
	}
	z01.PrintRune('\n')

	for i := 1; i < len(params); i++ {
		str := []rune(params[i])
		for j := 0; j < len(str); j++ {
			z01.PrintRune(str[j])
		}
		z01.PrintRune('\n')
	}
}
