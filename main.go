package main

import (
	"os"

	"github.com/01-edu/z01"
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

func printLn(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printSlc(slc []string) {
	for _, s := range slc {
		for _, r := range s {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

func vowelMirror(s string) string {
	var contVowels []rune
	var count int
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	rnSlc := []rune(s)

	for i := 0; i < len(rnSlc); i++ {
		for j := 0; j < len(vowels); j++ {
			if rnSlc[i] == vowels[j] || rnSlc[i] == vowels[j]-32 {
				contVowels = append(contVowels, rnSlc[i])
			}
		}
	}

	for i := 0; i < len(rnSlc); i++ {
		for j := 0; j < len(contVowels); j++ {
			if contVowels[j] == rnSlc[i] {
				if count == j {
					rnSlc = append(append(rnSlc[:i], contVowels[(0-j)+(len(contVowels)-1)]), rnSlc[i+1:]...)
					count++
					i++
				}
			}
		}
	}
	return string(rnSlc)
}
