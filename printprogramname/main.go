package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	str := []rune(s)
	var index int = 0
	for i := 0; i < len(s); i++ {
		if str[i] == '/' {
			index = i
		}
	}
	for i := index + 1; i < len(s); i++ {
		z01.PrintRune(str[i])
	}
}

func main() {
	PrintStr(os.Args[0])
}
