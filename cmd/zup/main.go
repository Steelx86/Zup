package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/dlambda/zup/pkg/ops"
)

const (
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
