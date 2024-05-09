package main

import "github.com/01-edu/z01"

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
