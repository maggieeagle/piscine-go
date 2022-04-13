package piscine

func IsAlpha(s string) bool {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if !(str[i] >= 65 && str[i] <= 90) && !(str[i] >= 97 && str[i] <= 122) && !(str[i] >= 48 && str[i] <= 57) {
			return false
		}
	}
	return true
}
