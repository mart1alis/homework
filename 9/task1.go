package main

import (
	"fmt"
	"unicode"
)

func bash(text string) string {
	result := ""
	for _, char := range text {
		if unicode.IsUpper(char) {
			result += string(('Z' - (int(char) - 'A')))
		} else if unicode.IsLower(char) {
			result += string(('z' - (int(char) - 'a')))
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	var text string

	fmt.Print("Text: ")
	fmt.Scanln(&text)

	fmt.Println("Result: ", bash(text))
}
