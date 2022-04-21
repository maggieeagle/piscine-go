package piscine

func Map(f func(int) bool, a []int) []bool {
	res := make([]bool, 0)
	for _, elem := range a {
		res = append(res, f(elem))
	}
	return res
}
