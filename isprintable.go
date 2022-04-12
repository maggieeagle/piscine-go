package piscine

func IsPrintable(s string) bool {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 127 {
			return false
		}
	}
	return true
}
