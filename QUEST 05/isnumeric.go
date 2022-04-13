package piscine

func IsNumeric(s string) bool {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] < 48 || str[i] > 57 {
			return false
		}
	}
	return true
}
