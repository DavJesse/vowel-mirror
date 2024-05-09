package main

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
