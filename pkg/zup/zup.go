package zup

import (
	"fmt"
	"os"
	"strings"
	"encoding/hex"

	"github.com/dlambda/zup/internal/encryption"
)

const (
	KEY_SIZE = 32
)

func InitZup(name string) (string, error) {
	if strings.HasSuffix(name, ".zup") {
		os.Create(name)
	} else {
		os.Create(name + ".zup")
	}

	key, err := encryption.GenerateKey(KEY_SIZE)
	if err != nil {
		return "", err
	}

	readableKey := hex.EncodeToString(key)

	return readableKey, nil
}
/*
func OpenZup(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
}
*/
func GenerateZupKey() (string, error) {
	key, err := encryption.GenerateKey(KEY_SIZE)
	if err != nil {
		return "", err
	}

	readableKey := hex.EncodeToString(key)

	return readableKey, nil
}
