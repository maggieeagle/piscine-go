package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 || nb == 1 {
		return 1
	}

	prevRes := RecursiveFactorial(nb - 1)

	if nb*prevRes > prevRes {
		return nb * prevRes
	}

	return 0
}
