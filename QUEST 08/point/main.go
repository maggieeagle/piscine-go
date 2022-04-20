package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func convert(num int) string {
	res := ""

	for num > 0 {
		res = string(num%10+48) + res
		num /= 10
	}
	return res
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := new(point)

	setPoint(points)

	printStr("x = " + convert(points.x) + ", y = " + convert(points.y))
	z01.PrintRune('\n')
}
