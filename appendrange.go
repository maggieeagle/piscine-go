package piscine

func AppendRange(min, max int) []int {
	arr := make([]int, 0)
	if min > max {
		return arr
	}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
