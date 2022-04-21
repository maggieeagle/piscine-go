package piscine

func Compare(a, b string) int {
	str1 := []byte(a)
	str2 := []byte(b)
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] < str2[i] {
			return -1
		}
		if str1[i] > str2[i] {
			return 1
		}
	}
	return 0
}

func SortWordArr(a []string) {
	params := a

	for i := 0; i < len(params)-1; i++ {
		for j := i + 1; j < len(params); j++ {
			if Compare(params[i], params[j]) == 1 {
				tmp := params[i]
				params[i] = params[j]
				params[j] = tmp
			}
		}
	}
}
