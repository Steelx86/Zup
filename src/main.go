package main

import (
	"fmt"
	"os"
)

func main() {
	fileLocation := os.Args[1]

	file, err := os.Open(fileLocation)

	if err != nil {
		fmt.Print(err)
		return
	}
	
	defer file.Close()
}