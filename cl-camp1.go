package main

import (
	"fmt"
	"os"
)

func main() {
	text := "ls -A -m -c -F"
	file, err := os.Create("mastertheLS")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)
}
