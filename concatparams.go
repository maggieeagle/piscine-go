package piscine

func ConcatParams(args []string) string {
	res := ""

	for i := 0; i < len(args); i++ {
		res += args[i]
		if i != len(args)-1 {
			res += "\n"
		}
	}
	return res
}
