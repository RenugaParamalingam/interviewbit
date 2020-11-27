package main

import "fmt"

func main() {
	fmt.Println(LengthOfLastWord("d"))
}

func LengthOfLastWord(A string) int {
	length := 0

	runes := []rune(A)

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]

		if string(r) == " " && length > 0 {
			return length
		} else if string(r) == " " && length == 0 {
			continue
		}

		length++
	}

	return length
}
