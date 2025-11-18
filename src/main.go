package main

import (
	"fmt"
	"os"
	"crypto/rand"
)

func main() {
	command := os.Args[1]

	if command == "open" {
		fileDir := os.Args[2]
		key := os.Args[3]

		file, err := os.Open(fileDir); if err != nil {
			fmt.Errorf("File failed to open: ", err)
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

func genKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := rand.Read(key); if err != nil {
		return nil, err
	}

	return key, err
}