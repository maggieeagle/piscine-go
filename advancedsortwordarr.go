package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
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
