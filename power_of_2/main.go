package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	num := "147573952589676412928"

	fmt.Println(powerOfTwo(num))
}

func powerOfTwo(A string) int {
	n := big.NewInt(0)
	if _, ok := n.SetString(A, 10); !ok {
		fmt.Printf("error parsing line %#v\n", A)
		os.Exit(1)
	}

	num := n.Int64()

	if num == 0 || num == 2 {
		return 1
	} else if num == 1 {
		return 0
	}

	for i := num; i > 1; {
		if i%2 != 0 {
			return 0
		}
		i = i / 2
	}

	return 1
}
