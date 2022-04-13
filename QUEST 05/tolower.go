package piscine

func ToLower(s string) string {
	str := []byte(s)
	var res string = ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 65 && str[i] <= 90 {
			res += string(str[i] + 32)
		} else {
			res += string(str[i])
		}
	}
	return res
}
