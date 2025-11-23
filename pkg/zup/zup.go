package zup

import (
	"encoding/hex"
	"os"
	"strings"

	"github.com/dlambda/zup/internal"
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

	key, err := internal.GenerateKey(KEY_SIZE)
	if err != nil {
		return "", err
	}

	readableKey := hex.EncodeToString(key)

	return readableKey, nil
}

func OpenZup(name string, key string) (internal.Zup, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return internal.Zup{}, err
	}

	encryptedContent := string(file)
	content, err := internal.Decrypt(encryptedContent, []byte(key))
	if err != nil {
		return internal.Zup{}, err
	}

	zupData, err := internal.ParseZup(content)
	if err != nil {
		return internal.Zup{}, err
	}

	return zupData, nil
}

func GenerateZupKey() (string, error) {
	key, err := internal.GenerateKey(KEY_SIZE)
	if err != nil {
		return "", err
	}

	readableKey := hex.EncodeToString(key)

	return readableKey, nil
}
