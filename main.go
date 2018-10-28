package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Fprint(os.Stderr, "Invalid arguments")
	}

	fmt.Printf("First file: %s\n", args[0]) // DEBUG
	fmt.Printf("Operation: %s\n", args[1])  // DEBUG
	fmt.Printf("Options: %s\n", args[2])    // DEBUG

	imgseq, err := FromString(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

}
