package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Compare(a, b string) int {
	str1 := []byte(a)
	str2 := []byte(b)
	if str1[0] == str2[0] {
		return 0
	}
	if str1[0] < str2[0] {
		return -1
	}
	return 1
}

func main() {
	params := os.Args

	for i := 1; i < len(params)-1; i++ {
		for j := i + 1; j < len(params); j++ {
			if Compare(params[i], params[j]) == 1 {
				tmp := params[i]
				params[i] = params[j]
				params[j] = tmp
			}
		}
	}

	for i := 1; i < len(params); i++ {
		str := []rune(params[i])
		for j := 0; j < len(str); j++ {
			z01.PrintRune(str[j])
		}
		z01.PrintRune('\n')
	}
}
