package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/dlambda/zup/pkg/encryption"
)

const (
	KEY_SIZE = 32
	helpMsg  = `Usage: zup [FILE] [KEY] 
Try 'zup help' for more information\n`
)

func main() {
	command := os.Args[1]

	switch command {
	case "-n":
		newZup(os.Args[2])
	case "-h":
		fmt.Println(helpMsg)
	case "-g":
		generateZupKey()
	case "-s": // sync with host
		return 
	case "-f": // forward port for hosting
		return
	default:
		fmt.Println(helpMsg)
	}
}

func newZup(name string) {
	if strings.HasSuffix(name, ".zup") {
		os.Create(name)
	} else {
		os.Create(name + ".zup")
	}

	key, err := encryption.generateKey(KEY_SIZE)
	if err != nil {
		fmt.Printf("Key generation error: %v\n", err)
	}

	readableKey := hex.EncodeToString(key)

	fmt.Printf("The key for your new zup file is: %s\n", readableKey)
}

func generateZupKey() {
	key, err := generateKey(KEY_SIZE)
	if err != nil {
		fmt.Printf("Key generation error: %v\n", err)
		return
	}

	readableKey := hex.EncodeToString(key)

	fmt.Printf("Generated key: %s\n", readableKey)
}
