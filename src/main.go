package main

import (
	"fmt"
	"os"
	"crypto/rand"
)

func main() {
	command := os.Args[1]

	switch command {
		case "open": openZup()
		case "new": newZup()
		case "help": fmt.Println(helpMsg)
		default: fmt.Println(helpMsg)
	}

}

func genKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := rand.Read(key); if err != nil {
		return nil, err
	}

	return key, err
}