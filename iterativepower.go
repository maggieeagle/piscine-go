package piscine

func IterativePower(nb int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res *= nb
	}
	return res
}
