package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := []int{124, 4, -3, 20, 124, 125}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
