package main

import (
	"os"
)

func main() {
	var mirrored string
	var argStr string

	args := os.Args[1:]

	if len(args) < 1 {
		printLn("")
		return
	}

	for _, str := range args {
		argStr += str + " "
	}

	mirrored = vowelMirror(argStr[:len(argStr)-1])
	printLn(mirrored)
}