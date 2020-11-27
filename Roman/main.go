package main

import "fmt"

func main() {
	fmt.Println("Number: ", RomanToInteger("IX"))
	fmt.Println(IntegerToRoman(3999))
}

func RomanToInteger(roman string) int {
	result := 0
	for i := 0; i < len(roman); i++ {
		r1 := value[string(roman[i])]

		if i+1 < len(roman) {
			r2 := value[string(roman[i+1])]
			if r1 >= r2 {
				result += r1
				continue
			}
			result += r2 - r1
			i++
		} else {
			result += r1
		}
	}
	return result
}

var value = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var i = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var x = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var c = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var m = []string{"", "M", "MM", "MMM"}

func IntegerToRoman(num int) string {
	thousands := m[num/1000]
	hundreds := c[num%1000/100]
	tens := x[num%100/10]
	ones := i[num%10]

	return thousands + hundreds + tens + ones
}
