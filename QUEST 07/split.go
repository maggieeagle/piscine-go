package piscine

func getSubstring(s string, f, l int) string {
	res := ""
	str := []rune(s)
	for i := f; i < l; i++ {
		res += string(str[i])
	}
	return res
}

func Split(s, sep string) []string {
	var res []string

	str := []rune(s)

	subs := ""
	for i := 0; i < len(str)-len(sep)+1; i++ {
		if getSubstring(s, i, i+len(sep)) != sep {
			subs += string(str[i])
		} else {
			if len(subs) != 0 {
				res = append(res, subs)
				i += len(sep) - 1
			}
			subs = ""
		}
	}
	for i := len(str) - len(sep) + 1; i < len(str); i++ {
		subs += string(str[i])
	}
	res = append(res, subs)
	return res
}
