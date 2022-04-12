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

func Sort(table []byte) []byte {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] < table[j] {
				var tmp byte = table[i]
				table[i] = table[j]
				table[j] = tmp
			}
		}
	}
	return table
}

func PrintNbrInOrder(n int) {
	if n >= 0 {
		str = ""

		if n == 0 {
			AddToString(0)
		} else {
			for n > 0 {
				AddToString(n % 10)
				n /= 10
			}
		}

		s := []byte(str)
		sort := Sort(s)

		for i := len(sort) - 1; i > -1; i-- {
			z01.PrintRune(rune(sort[i]))
		}
	}
}
