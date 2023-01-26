package roman

import (
	"fmt"
	"strings"
)

// Holds a unit of a Roman numeral together with its integer value
type numeral struct {
	roman string
	value int
}

// MARK: Roman to Integer

// Checks for correct order, e.g. "I" can precede "V", but "I" cannot precede "X"
func hasValidRankDifference(lhs numeral, rhs numeral) bool {
	numeralsOrder := "MDCLXVI"
	/*
		Correct and incorrect numeral placement reference:
				VI	XI	IV	IX
		ranks	5,6	4,6	6,5	6,4
		allowed	✓	✓	✓	⨯

		So, `left` cannot be 2 levels higher than `right`
	*/
	leftRank := strings.Index(numeralsOrder, lhs.roman)
	rightRank := strings.Index(numeralsOrder, rhs.roman)
	return leftRank <= rightRank || leftRank-rightRank < 2
}

// A map of primitives needed to decypher the roman numerals
var numeralPrimitives = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// Converts a single primitive Roman numeral into an integer
func singleRomanNumeralToInteger(char string) (int, error) {
	val, ok := numeralPrimitives[char]
	if ok {
		return val, nil
	} else {
		return 0, fmt.Errorf("%v is not a Roman numeral letter", char)
	}
}

// Converts a Roman numeral into Arabic, integer form
func RomanToInteger(roman string) (int, error) {
	if roman == "" {
		return 0, fmt.Errorf("empty input string")
	}

	// Convert each letter of Roman numerals into a letter-value pair
	var expandedValues = []numeral{}
	for _, char := range strings.ToUpper(roman) {
		var value, error = singleRomanNumeralToInteger(string(char))
		if error != nil {
			return 0, error
		}
		expandedValues = append(expandedValues, numeral{string(char), value})
	}

	// Take values one by one:
	// - If current value of symbol is greater than or equal to the value of next symbol, then add this value to the running total.
	// - else subtract this value by adding the value of next symbol to the running total.
	var total = 0

	for i := 0; i < len(expandedValues); i++ {
		// Current and next, if there is next
		var current = expandedValues[i]
		var next numeral
		if i+1 < len(expandedValues) {
			next = expandedValues[i+1]
		}

		if next.value != 0 {
			if !hasValidRankDifference(current, next) {
				return 0, fmt.Errorf("not a valid Roman numeral")
			}
			if current.value >= next.value {
				// If `current` is larger than `next` or if it's the last one, add its value
				// Also, make sure that rakning is correct (e.g. "IV" is good, "IX" is not)
				total += current.value
			} else {
				// If after the `current` there is a larger value, like in "IV", then add their differnce
				total += next.value - current.value
				i++
			}
		} else {
			// If there is no `next`, just add what's there
			total += current.value
		}
	}

	return total, nil
}

// MARK: Integer to Roman

// Ordered list of Roman numerals, needed for quickly composing them from an integer
var numerals = []numeral{
	{roman: "M", value: 1000},
	{roman: "CM", value: 900},
	{roman: "D", value: 500},
	{roman: "CD", value: 400},
	{roman: "C", value: 100},
	{roman: "XC", value: 90},
	{roman: "L", value: 50},
	{roman: "XL", value: 40},
	{roman: "X", value: 10},
	{roman: "IX", value: 9},
	{roman: "V", value: 5},
	{roman: "IV", value: 4},
	{roman: "I", value: 1},
}

// Converts an integer number to a Roman numeral string
func IntegerToRoman(number int) (string, error) {

	if number < 0 {
		return "", fmt.Errorf("cannot convert negative numbers to Roman numeral notation")
	}
	if number == 0 {
		return "", fmt.Errorf("0 cannot be represented in Roman numeral notation")
	}
	if number > 3999 {
		return "", fmt.Errorf("numbers greater than 3999 cannot be represented in Roman numeral notation")
	}

	var output = ""

	for _, numeral := range numerals {
		for number >= numeral.value {
			output += numeral.roman
			number -= numeral.value
		}
	}

	return output, nil
}
