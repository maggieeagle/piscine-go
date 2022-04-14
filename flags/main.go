package main

import (
	"fmt"
	"os"
)

var (
	isInsert     bool   = false
	insertString string = ""
	isOrder      bool   = false
	isHelp       bool   = false
)

var getString string = ""

func GetSubstring(s []rune, f, l int) string {
	res := ""
	for i := f; i < l; i++ {
		res += string(s[i])
	}
	return res
}

func CheckParam(s string) {
	p := []rune(s)

	if p[0] == '-' {
		switch p[1] {
		case 'i':
			isInsert = true
			insertString = GetSubstring(p, 2, len(p))
		case 'o':
			isOrder = true
		case 'h':
			isHelp = true
		case '-':
			switch p[2] {
			case 'i':
				isInsert = true
				insertString = GetSubstring(p, 9, len(p))
			case 'o':
				isOrder = true
			case 'h':
				isHelp = true
			}
		}
	} else {
		getString = s
	}
}

func SortString(s string) string {
	str := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if str[i] > str[j] {
				tmp := str[i]
				str[i] = str[j]
				str[j] = tmp
			}
		}
	}

	var res string = ""

	for i := 0; i < len(s); i++ {
		res += string(str[i])
	}
	return res
}

func main() {
	params := os.Args

	if len(params) == 1 {
		isHelp = true
	} else {
		for i := 1; i < len(params); i++ {
			CheckParam(params[i])
		}
	}

	if isInsert {
		getString += insertString
	}
	if isOrder {
		fmt.Println(SortString(getString))
	}
	if isHelp {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("    This flag inserts the string passed as argument.")

		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("    This flag will behave like a boolean, if it is called it will order the argument.")
	}
}
