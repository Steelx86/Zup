package main

import (
	
)

func main() {

}

func openZup(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
}
