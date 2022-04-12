package piscine

func IsLetterHigh(l byte) bool {
	if l >= 65 && l <= 90 {
		return true
	}
	return false
}

func IsLetterLow(l byte) bool {
	if l >= 97 && l <= 122 {
		return true
	}
	return false
}

func LowLetterToHigh(l byte) byte {
	return l - 32
}

func HighLetterToLow(l byte) byte {
	return l + 32
}

func Capitalize(s string) string {
	str := []byte(s)
	var res string = ""
	if IsLetterLow(str[0]) {
		res += string(LowLetterToHigh(str[0]))
	} else {
		res += string(str[0])
	}
	for i := 1; i < len(str); i++ {
		if (IsLetterHigh(str[i-1]) || IsLetterLow(str[i-1])) && IsLetterHigh(str[i]) {
			res += string(HighLetterToLow(str[i]))
		} else if !(IsLetterHigh(str[i-1]) || IsLetterLow(str[i-1])) && IsLetterLow(str[i]) {
			res += string(LowLetterToHigh(str[i]))
		} else {
			res += string(str[i])
		}
	}
	return res
}
