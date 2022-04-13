package piscine

func getDigits(s string) string {
	str := []byte(s)
	var res string = ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 48 && str[i] <= 57 || len(res) == 0 && str[i] == 45 {
			res += string(str[i])
		}
	}
	return res
}

func TrimAtoi(s string) int {
	return Atoi(getDigits(s))
}
