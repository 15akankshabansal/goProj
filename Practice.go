package main

import (
	"fmt"
)

func reverseWords(s string) string {
	words := string.Fields(s) // Split the string into words
	for i, word := range words {
		if len(word) >= 5 {
			// Reverse the word if it has five or more letters
			words[i] = rev(word)
		}
	}
	return string.Join(words, " ") // Join the words back into a string
}
func rev(s string) string {

	charstr := []rune(s)
	for _, char := range charstr {
		fmt.Printf("%c", char)
	}
	for i, j := 0, len(charstr)-1; i < j; i, j = i+1, j-1 {
		charstr[i], charstr[j] = charstr[j], charstr[i]
	}
	return string(charstr)
}

func main() {
	str := "trfyjhl"
	reversedValue := rev(str)
	reverseWords(str)

	fmt.Print("My reversed string is", reversedValue)
}
