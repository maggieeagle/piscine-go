package piscine

func StrRev(s string) string {
	str1 := []rune(s)
	var str2 string = ""

	for i := len(str1) - 1; i > -1; i-- {
		str2 += string(str1[i])
	}
	return str2
}
