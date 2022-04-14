package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Pow(a int, b int) int {
	pow := 1
	if b == 0 {
		return 1
	}
	for i := 0; i < b; i++ {
		pow *= a
	}
	return pow
}

func Reverse(s string) string {
	str1 := []rune(s)
	var str2 string = ""

	for i := len(str1) - 1; i > -1; i-- {
		str2 += string(str1[i])
	}
	return str2
}

func RuneToIntWithEx2(digit rune) int {
	switch digit {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case '+':
		return -2
	case '-':
		return -3
	default:
		return -1
	}
}

func Atoi(s string) int {
	str := []rune(Reverse(s))

	var n int = 0

	var isNegative int = 1
	for i := 0; i < len(str); i++ {
		var digit int = RuneToIntWithEx2(str[i])
		if i == len(str)-1 {
			if digit == -2 {
				break
			}
			if digit == -3 {
				isNegative = -1
				break
			}
		}
		if digit == -1 || digit == -2 || digit == -3 {
			return 0
		}
		n += Pow(10, i) * digit
	}
	return isNegative * n
}

func PrintStr(s string) {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		z01.PrintRune(str[i])
	}
}

func PrintUp(mas []string) {
	for i := 2; i < len(mas); i++ {
		num := Atoi(mas[i])
		var str string = " "
		if num >= 1 && num <= 26 {
			str = string(64 + num)
		}
		PrintStr(str)
	}
}

func PrintDown(mas []string) {
	for i := 1; i < len(mas); i++ {
		num := Atoi(mas[i])
		var str string = " "
		if num >= 1 && num <= 26 {
			str = string(96 + num)
		}
		PrintStr(str)
	}
}

func main() {
	params := os.Args
	var isUp bool = false

	if len(params) > 1 {

		if params[1] == "--upper" {
			isUp = true
		}

		if isUp {
			PrintUp(params)
		} else {
			PrintDown(params)
		}

		z01.PrintRune('\n')
	}
}
