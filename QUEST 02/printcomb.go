package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	digits := []rune("0123456789")
	for i := 0; i < len(digits)-2; i++ {
		for j := i + 1; j < len(digits)-1; j++ {
			for k := j + 1; k < len(digits); k++ {
				z01.PrintRune(digits[i])
				z01.PrintRune(digits[j])
				z01.PrintRune(digits[k])
				if !(i == 7 && j == 8 && k == 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
