package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				var tmp int = table[i]
				table[i] = table[j]
				table[j] = tmp
			}
		}
	}
}
