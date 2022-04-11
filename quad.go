package piscine

import (
	"github.com/01-edu/z01"
)

func PrintQuad(x, y int, s1, s2, s3, s4, s5, s6 rune) {
	if y > 0 {
		if x > 0 {
			z01.PrintRune(s1)
		}
		for i := 1; i < x-1; i++ {
			z01.PrintRune(s2)
		}
		if x > 1 {
			z01.PrintRune(s3)
		}
		z01.PrintRune('\n')
	}
	for j := 1; j < y-1; j++ {
		if x > 0 {
			z01.PrintRune(s6)
		}
		for i := 1; i < x-1; i++ {
			z01.PrintRune(' ')
		}
		if x > 1 {
			z01.PrintRune(s6)
		}
		z01.PrintRune('\n')
	}
	if y > 1 {
		if x > 0 {
			z01.PrintRune(s4)
		}
		for i := 1; i < x-1; i++ {
			z01.PrintRune(s2)
		}
		if x > 1 {
			z01.PrintRune(s5)
		}
		z01.PrintRune('\n')
	}
}

// s1-s2-s3
// s6	s6
// s4-s2-s5

func QuadA(x, y int) {
	s1 := 'o'
	s2 := '-'
	s3 := 'o'
	s4 := 'o'
	s5 := 'o'
	s6 := '|'
	PrintQuad(x, y, s1, s2, s3, s4, s5, s6)
}

func QuadB(x, y int) {
	s1 := '/'
	s2 := '*'
	s3 := '\\'
	s4 := '\\'
	s5 := '/'
	s6 := '*'
	PrintQuad(x, y, s1, s2, s3, s4, s5, s6)
}

func QuadC(x, y int) {
	s1 := 'A'
	s2 := 'B'
	s3 := 'A'
	s4 := 'C'
	s5 := 'C'
	s6 := 'B'
	PrintQuad(x, y, s1, s2, s3, s4, s5, s6)
}

func QuadD(x, y int) {
	s1 := 'A'
	s2 := 'B'
	s3 := 'C'
	s4 := 'A'
	s5 := 'C'
	s6 := 'B'
	PrintQuad(x, y, s1, s2, s3, s4, s5, s6)
}

func QuadE(x, y int) {
	s1 := 'A'
	s2 := 'B'
	s3 := 'C'
	s4 := 'C'
	s5 := 'A'
	s6 := 'B'
	PrintQuad(x, y, s1, s2, s3, s4, s5, s6)
}
