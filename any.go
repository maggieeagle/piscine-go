package piscine

func Any(f func(string) bool, a []string) bool {
	for _, elem := range a {
		if f(elem) {
			return true
		}
	}
	return false
}
