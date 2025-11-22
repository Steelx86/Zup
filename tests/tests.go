package tests

import (
	"os"
	"testing"
)

func testFunctionality(t *testing.T) {
	files, err := os.ReadDir("entries")
	if err != nil {
		return nil, err
	}
}

