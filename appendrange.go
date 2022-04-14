package piscine

func CreateArrayOfSize(size int) []int {
	var answer []int
	for i := 0; i < size; i++ {
		answer = append(answer, i+1)
	}
	return answer
}

func AppendRange(min, max int) []int {
	var size int = 0

	if min < max {
		size = max - min
	}

	arr := CreateArrayOfSize(size)

	if size > 0 {
		for i := min; i < max; i++ {
			arr = append(arr, i)
		}
	}
	return arr
}
