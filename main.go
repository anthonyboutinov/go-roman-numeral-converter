package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/anthonyboutinov/go-roman-numerals-converter/roman"
)

func integerToRoman(num int) {
	roman, error := roman.IntegerToRoman(num)
	if error != nil {
		fmt.Printf("Unable to convert integer to Roman: %v\n", error)
		return
	}
	fmt.Println("Roman numeral:", roman)
}

func romanToInteger(_roman string) {
	integer, error := roman.RomanToInteger(_roman)
	if error != nil {
		fmt.Printf("Unable to convert Roman to integer: %v\n", error)
		return
	}
	fmt.Println("Number:", integer)
}

func main() {

	// Prompt the user for input
	fmt.Println("Enter a number to convert to Roman numeral (or a Roman numeral to convert to number):")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		// Check if the input is a number or a Roman numeral
		// Atoi converts input into an Int number, or throws an error
		num, error := strconv.Atoi(input)

		// Call the appropriate conversion function
		if error == nil {
			// If the input is number, convert to Roman numeral
			integerToRoman(num)
		} else {
			// Check if it's a `float`
			_, error := strconv.ParseFloat(input, 64)
			if error == nil {
				fmt.Println("Floating point numbers cannot be represented in Roman numeral notation")
			} else {
				// If the input is a string, convert it to integer form
				romanToInteger(input)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
