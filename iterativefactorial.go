package piscine

func IterativeFactorial(nb int) int {

	if nb < 1 {
		return 0
	}

	var result int = 1
	for i := 1; i <= nb; i++ {
		result *= i
	}

	return result
}
