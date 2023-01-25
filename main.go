package main

import (
	"fmt"

	"github.com/anthonyboutinov/go-roman-numerals-converter/roman"
)

func main() {
	input := "MMDLXXXII"
	val, err := roman.RomanToInteger(input)
	fmt.Printf(`%v = %v, err: %v`, input, val, err)
}
