package piscine

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
