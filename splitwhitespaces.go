package piscine

var separators []rune = []rune{' ', '\n', '\t'}

func IsSeparator(r rune) bool {
	for i := 0; i < len(separators); i++ {
		if r == separators[i] {
			return true
		}
	}
	return false
}

func SplitWhiteSpaces(s string) []string {
	var res []string

	str := []rune(s)

	subs := ""
	for i := 0; i < len(str); i++ {
		if !IsSeparator(str[i]) {
			subs += string(str[i])
		} else {
			if len(subs) != 0 {
				res = append(res, subs)
			}
			subs = ""
		}
		if i == len(str)-1 && subs != "" {
			res = append(res, subs)
		}
	}
	return res
}
