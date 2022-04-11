package piscine

func IterativeFactorial(nb int) int {
	if nb < 1 {
		return 0
	}

	var result int = 1
	for i := 1; i <= nb; i++ {
		oldResult := result
		result *= i
		if result < oldResult {
			return 0
		}
	}

	return result
}
