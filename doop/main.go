package main

import (
	"os"
)

var (
	MIN_INT = []rune("-9223372036854775808")
	MAX_INT = []rune("9223372036854775807")
)

func isNumeric(s string) bool {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if (str[i] < 48 || str[i] > 57) && str[i] != 45 {
			return false
		}
	}
	return true
}

func isSign(s string) bool {
	str := []byte(s)
	if len(s) > 1 {
		return false
	}
	signs := [5]rune{'+', '-', '/', '*', '%'}
	for _, sign := range signs {
		if rune(str[0]) == sign {
			return true
		}
	}
	return false
}

func isInt(s string) bool {
	str := []rune(s)
	if !isNumeric(s) || (len(s) > 19 && str[0] != '-') || (len(s) > 20 && str[0] == '-') {
		return false
	}
	if str[0] == '-' && len(str) == len(MIN_INT) {
		for i := 1; i < len(str); i++ {
			if int(str[i]) > int(MIN_INT[i]) {
				return false
			}
		}
	} else if len(str) == len(MAX_INT) {
		for i := 0; i < len(str); i++ {
			if str[i] > MAX_INT[i] {
				return false
			}
		}
	}
	return true
}

func makeOperation(a, b int, sign rune) []int {
	res := []int{0, 0}
	switch sign {
	case '+':
		tmp := a
		old_tmp := a
		for i := 0; i < b; i++ {
			tmp += 1
			if tmp < old_tmp {
				res[1] = 1
			}
		}
		res[0] = tmp
	case '-':
		tmp := a
		old_tmp := a
		for i := 0; i < b; i++ {
			tmp -= 1
			if tmp > old_tmp {
				println("Caraul!")
				res[1] = 1
			}
		}
		res[0] = tmp
	case '/':
		res[0] = a / b
	case '*':
		mult := 1
		if a > 0 && b < 0 {
			b *= -1
			mult = -1
		}
		if a < 0 && b > 0 {
			a *= -1
			mult = -1
		}
		if a < 0 && b < 0 {
			a *= -1
			b *= -1
		}
		tmp := a
		old_tmp := a
		for i := 0; i < b-1; i++ {
			tmp += a
			if tmp < old_tmp {
				res[1] = 1
			}
		}
		res[0] = mult * tmp
	case '%':
		res[0] = a % b
	}
	return res
}

func main() {
	params := os.Args

	if len(params) == 4 && isInt(params[1]) && isSign(params[2]) && isInt(params[3]) {
		a := Atoi(params[1])
		b := Atoi(params[3])
		sign := []rune(params[2])[0]

		if sign == '/' && b == 0 {
			print("No division by 0")
		} else if sign == '%' && b == 0 {
			print("No modulo by 0")
		} else {
			res := makeOperation(a, b, sign)
			if res[1] == 0 {
				println(res[0])
			}
		}
	}
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
