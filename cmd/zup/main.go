package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
	case "open":
		openZup(os.Args[2])
	case "new":
		newZup(os.Args[2])
	case "help":
		fmt.Println(helpMsg)
	case "generate":
		generateZupKey()
	case "sync": // sync with host
		return // place holder
	case "host": // open port for hosting
		return
	case "":
		REPL()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println(helpMsg)
	}
}
