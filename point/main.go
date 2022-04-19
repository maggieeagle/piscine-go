package main

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := new(point)

	setPoint(points)

	print("x = ")
	print(points.x)
	print(", y = ")
	print(points.y)
	println()
}
