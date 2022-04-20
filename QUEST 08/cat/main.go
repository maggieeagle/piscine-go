package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func readFromFile(f string) {
	file, err := os.Open(f)
	if err != nil {
		printStr("ERROR: " + err.Error())
		z01.PrintRune('\n')
		os.Exit(1)
	}
	dat, _ := file.Stat()
	content := make([]byte, dat.Size())
	file.Read(content)
	printStr(string(content))
	file.Close()
}

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for i := 1; i < len(os.Args); i++ {
			readFromFile(os.Args[i])
		}
	}
}
