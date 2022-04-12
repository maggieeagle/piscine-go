package piscine

func getSubstring(s string, f, l int) string {
	res := ""
	str := []rune(s)
	for i := f; i < l; i++ {
		res += string(str[i])
	}
	return res
}

func Index(s string, toFind string) int {
	str := []rune(s)
	toF := []rune(toFind)
	for i := 0; i < len(str)-len(toF); i++ {
		sub := getSubstring(s, i, i+len(toF))
		if Compare(sub, toFind) == 0 {
			return i
		}
	}
	return -1
}
