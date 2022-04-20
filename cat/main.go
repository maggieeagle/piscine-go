package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func main() {
	switch len(os.Args) {
	case 1:
		io.Copy(os.Stdout, os.Stdin)

	case 2:
		dat, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			printStr("ERROR: " + err.Error())
			z01.PrintRune('\n')
			os.Exit(1)
		} else {
			printStr(string(dat))
		}
	case 3:
		dat, err := ioutil.ReadFile(os.Args[1])
		dat2, err2 := ioutil.ReadFile(os.Args[2])
		if err != nil {
			printStr("ERROR: " + err.Error())
			z01.PrintRune('\n')
		} else {
			printStr(string(dat))
		}
		if err2 != nil {
			printStr("ERROR: " + err.Error())
			z01.PrintRune('\n')
		} else {
			printStr(string(dat2))
		}
	}
}
