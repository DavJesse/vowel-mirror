package main

import (
	"os"
)

func main() {
	var mirrored string
	var argStr string

	//capture all arguments
	args := os.Args[1:]

	//exit program if no arguments are provided
	if len(args) < 1 {
		printLn("You need at least one argument to run this program")
		return
	}

	//concat arguments with spaces
	for _, str := range args {
		argStr += str + " "
	}

	//mirror vowels in the string
	mirrored = vowelMirror(argStr[:len(argStr)-1])
	//print result
	printLn(mirrored)
}
