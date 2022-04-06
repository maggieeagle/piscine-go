package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	digits := []rune("0123456789")

	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			for k := 0; k < len(digits); k++ {
				for l := j + 1; l < len(digits); l++ {
					if !(i == k && j == l) {
						z01.PrintRune(digits[i])
						z01.PrintRune(digits[j])
						z01.PrintRune(' ')
						z01.PrintRune(digits[k])
						z01.PrintRune(digits[l])

						if !(i == 9 && j == 9 && k == 9 && l == 8) {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
