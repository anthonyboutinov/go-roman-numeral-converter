package main

import (
	"fmt"
)

var numerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func singleRomanNumeralToInteger(char string) (int, error) {
	val, ok := numerals[char]
	if ok {
		return val, nil
	} else {
		return 0, fmt.Errorf("%v is not a roman numeral", char)
	}
}

func RomanToInteger(roman string) (int, error) {
	// Convert each symbol of Roman Numerals into the value it represents
	var rawValues = []int{}
	for _, char := range roman {
		var value, error = singleRomanNumeralToInteger(string(char))
		if error != nil {
			return 0, error
		}
		rawValues = append(rawValues, value)
	}

	// Take values one by one:
	// - If current value of symbol is greater than or equal to the value of next symbol, then add this value to the running total.
	// - else subtract this value by adding the value of next symbol to the running total.
	var total = 0
	for i := 0; i < len(rawValues); i++ {
		// Current and next, if there is next
		var current = rawValues[i]
		var next = 0
		if i+1 < len(rawValues) {
			next = rawValues[i+1]
		}

		if (next != 0 && current >= next) || next == 0 {
			// If current is larger than next or if it's the last one, add its value
			total += current
		} else {
			// If after the current there is a larger value, like in "IV", then add their differnce
			total += next - current
			i++
		}
	}

	return total, nil
}

func IntegerToRoman(integer int) (string, error) {
	return "", nil
}
