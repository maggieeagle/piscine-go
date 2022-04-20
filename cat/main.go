package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("Hello\nHello\n^C")
	case 2:
		dat, err := ioutil.ReadFile("../" + os.Args[1])
		if err != nil {
			fmt.Println("ERROR: open " + os.Args[1] + ": no such file or directory")
		}
		fmt.Print(string(dat))
	case 3:
		dat, err := ioutil.ReadFile("../" + os.Args[1])
		dat2, err2 := ioutil.ReadFile("../" + os.Args[2])
		if err != nil {
			fmt.Println("ERROR: open " + os.Args[1] + ": no such file or directory")
		}
		fmt.Print(string(dat))
		if err2 != nil {
			fmt.Println("ERROR: open " + os.Args[2] + ": no such file or directory")
		}
		fmt.Print(string(dat2))
	}
}
