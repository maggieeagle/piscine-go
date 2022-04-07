package piscine

func UltimateDivMod(a *int, b *int) {
	var a1 = *a
	*a = *a / *b
	*b = a1 - *b**a
}
