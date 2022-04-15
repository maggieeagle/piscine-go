package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string

	str := []rune(s)

	subs := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			subs += string(str[i])
		} else {
			res = append(res, subs)
			subs = ""
		}
		if i == len(str)-1 && subs != "" {
			res = append(res, subs)
		}
	}
	return res
}
