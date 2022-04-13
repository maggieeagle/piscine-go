package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(str[i])
	}
}

func main() {
	println(os.Args[0])
}
