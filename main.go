package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Fprint(os.Stderr, "Invalid arguments\n")
		os.Exit(1)
	}

	fmt.Printf("First file: %s\n", args[0]) // DEBUG
	fmt.Printf("Operation: %s\n", args[1])  // DEBUG
	fmt.Printf("Options: %s\n", args[2])    // DEBUG

	imgseq, err := initImgSeq(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	var op timeshift
	switch args[1] {
	case "timeshift":
		{
			op, err = initTimeshift(args[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
		}
	}
	op.Apply(imgseq)
	os.Exit(0)
}
