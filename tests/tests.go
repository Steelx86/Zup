package test

import (
	"testing"

	"github.com/dlambda/zup/pkg/src"
)

/* Encryption Tests */

func TestEncryption(t *testing.T) {
	key, err := src.GenerateKey(32)
	if err != nil {
		t.Fatalf("Key generation failed: %v", err)
	}
}

/* Network Tests */

func TestNetwork(t *testing.T) {

}

/* Command Handling Tests */

func TestCommandHandler(t *testing.T) {

}

/* REPL Tests */

func TestREPL(t *testing.T) {

}
