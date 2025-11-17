package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]

	if command == "open" {
		fileDir := os.Args[2]
		key := os.Args[3]

		file, err := os.Open(fileDir)

		if err != nil {
			fmt.Errorf("File failed to open: %v", err)
			return
		}

		defer file.Close()
	} else if command == "new" {
		fileName := os.Args[2]

		os.Create(fileName)
	} else if command == "help" {
		fmt.Print(helpMsg)
	} else {
		fmt.Print(helpMsg)
	}

}