package piscine

func ToUpper(s string) string {
	str := []byte(s)
	var res string = ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 97 && str[i] <= 122 {
			res += string(str[i] - 32)
		} else {
			res += string(str[i])
		}
	}
	return res
}
