package piscine

func IsUpper(s string) bool {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] < 65 || str[i] > 90 {
			return false
		}
	}
	return true
}
