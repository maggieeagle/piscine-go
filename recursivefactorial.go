package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb <= 1 {
		return 1
	}

	res := nb * RecursiveFactorial(nb-1)

	if res < 0 {
		return 0
	} else {
		return res
	}
}
