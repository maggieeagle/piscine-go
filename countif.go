package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int = 0
	for _, elem := range tab {
		if f(elem) {
			count++
		}
	}
	return count
}
