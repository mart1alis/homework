package main

import (
	"fmt"
	"unicode"
)

func caesar(text string, shift int) string {
	result := ""
	for _, char := range text {
		if unicode.IsUpper(char) {
			result += string((int(char)-'A'+shift)%33 + 'A')
		} else if unicode.IsLower(char) {
			result += string((int(char)-'a'+shift)%33 + 'a')
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	var text string
	var shift int

	fmt.Print("Text: ")
	fmt.Scanln(&text)

	fmt.Print("Shift: ")
	fmt.Scanln(&shift)

	fmt.Println("Result:", caesar(text, shift))
}
