package piscine

func MakeRange(min, max int) []int {
	if min < max {

		arr := make([]int, max-min)

		count := min
		for i := 0; i < len(arr); i++ {
			arr[i] = count
			count++
		}
		return arr
	}
	arr := make([]int, 0)
	return arr
}
