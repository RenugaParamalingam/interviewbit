package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(IsPalindromic("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindromic("race a car"))
}

func IsPalindromic(value string) bool {
	value = sanitize(value)
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}
	return true
}

func sanitize(value string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	safe := reg.ReplaceAllString(value, "")

	return strings.ToLower(strings.Trim(safe, ""))
}

func Palindrome() {
	originalStr := "madam"
	revStr := ""

	for i := len(originalStr) - 1; i >= 0; i-- {
		revStr += string(originalStr[i])
	}

	if strings.EqualFold(originalStr, revStr) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}
