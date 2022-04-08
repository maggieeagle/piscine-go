package piscine

import (
	"github.com/01-edu/z01"
)

var str string = ""

const MIN_INT_64 = -9223372036854775808

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

// некоторые числа из тестов не помещаются в int n
func PrintNbr(n int) {

	str = ""

	if n == MIN_INT_64 {
		str = "-9223372036854775808"

		s := []rune(str)

		for i := 0; i < len(s); i++ {
			z01.PrintRune(s[i])
		}
	} else {

		if n < 0 {
			z01.PrintRune('-')
			n *= -1
		}

		if n == 0 {
			AddToString(0)
		} else {
			for n > 0 {
				AddToString(n % 10)
				n /= 10
			}
		}

		s := []rune(str)

		for i := len(s) - 1; i > -1; i-- {
			z01.PrintRune(s[i])
		}
	}
}
