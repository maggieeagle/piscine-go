package main

import (
	"github.com/01-edu/z01"
)

func main() {
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
		'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
		't', 'u', 'v', 'w', 'x', 'y', 'z', '\n'}
	for i := 0; i < 27; i++ {
		z01.PrintRune(alphabet[i])
	}
}
