package piscine

func ForEach(f func(int), a []int) {
	for _, elem := range a {
		f(elem)
	}
}
