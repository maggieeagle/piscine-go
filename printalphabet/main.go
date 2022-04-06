package main

import (
	"github.com/01-edu/z01"
)

func main() {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz\n")
	for i := 0; i < 27; i++ {
		z01.PrintRune(alphabet[i])
	}
}
