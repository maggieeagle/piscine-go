package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	var r rune = 'F'
	if nb < 0 {
		r = 'T'
	}
	z01.PrintRune(r)
	z01.PrintRune('\n')
}
