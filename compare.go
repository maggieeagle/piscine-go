package piscine

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Compare(a, b string) int {
	str1 := []byte(a)
	str2 := []byte(b)

	for i := 0; i < min(len(str1), len(str2)); i++ {
		if str1[i] < str2[i] {
			return -1
		}
	}
	if len(str1) > len(str2) {
		return 1
	}
	return 0
}
