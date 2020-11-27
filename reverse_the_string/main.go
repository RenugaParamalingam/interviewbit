package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	A := " the sky is blue 1   ! "

	space := regexp.MustCompile(`\s+`)

	arr := strings.Split(A, " ")

	var rev []string
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			rev = append(rev, space.ReplaceAllString(arr[i], ""))
		}
	}

	fmt.Println(strings.Join(rev, " "))
}
