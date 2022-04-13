package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := []rune(os.Args[0])
	var index int = 0
	for i := 0; i < len(str); i++ {
		if str[i] == '/' {
			index = i
		}
	}
	for i := index + 1; i < len(str); i++ {
		z01.PrintRune(str[i])
	}
}
