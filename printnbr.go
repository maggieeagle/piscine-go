package piscine

import (
	"github.com/01-edu/z01"
)

var str string = ""

func AddToString(digit int) {
	switch digit {
	case 0:
		str += "0"
	case 1:
		str += "1"
	case 2:
		str += "2"
	case 3:
		str += "3"
	case 4:
		str += "4"
	case 5:
		str += "5"
	case 6:
		str += "6"
	case 7:
		str += "7"
	case 8:
		str += "8"
	case 9:
		str += "9"
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n += -2 * n
	}

	for n > 0 {
		AddToString(n % 10)
		n /= 10
	}

	s := []rune(str)

	for i := len(s) - 1; i > -1; i-- {
		z01.PrintRune(s[i])
	}

	z01.PrintRune('\n')
}
