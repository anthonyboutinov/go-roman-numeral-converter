package main

import "testing"

func TestRomanToIntegerMeaningfulInput(t *testing.T) {
	input := "CXXXIV"
	want := 134
	got, error := RomanToInteger(input)
	if got != want || error != nil {
		t.Fatalf(`RomanToInteger("%s") = %v, %v, want %v`, input, got, error, want)
	}
}

func TestRomanToIntegerMeaningfulInputLowerCase(t *testing.T) {
	input := "cxxXiv"
	want := 134
	got, error := RomanToInteger(input)
	if got != want || error != nil {
		t.Fatalf(`RomanToInteger("%s") = %v, %v, want %v`, input, got, error, want)
	}
}

func TestRomanToIntegerAtLeastOneOfEachRomanNumerals(t *testing.T) {
	input := "MDCLXVI"
	want := 1666
	got, error := RomanToInteger(input)
	if got != want || error != nil {
		t.Fatalf(`RomanToInteger("%s") = %v, %v, want %v`, input, got, error, want)
	}
}

func TestRomanToIntegerJibberishOfRomanNumerals(t *testing.T) {
	input := "ICXLXXIVMD"
	want := "to raise an error"
	got, error := RomanToInteger(input)
	if error == nil {
		t.Fatalf(`RomanToInteger("%s") = %v, %v, want %v`, input, got, error, want)
	}
}

func TestRomanToIntegerWithInvalidCharacters(t *testing.T) {
	input := "XI6V"
	want := "to raise a (non roman numeral) error"
	got, error := RomanToInteger(input)
	if error == nil {
		t.Fatalf(`RomanToInteger("%s") = %v, %v, want %v`, input, got, error, want)
	}
}

func TestIntegerToRomanMeaningfull(t *testing.T) {
	input := 628
	want := "DCXXVIII"
	got, error := IntegerToRoman(input)
	if got != want || error != nil {
		t.Fatalf(`IntegerToRoman("%v") = %v, %v, want %v`, input, got, error, want)
	}
}

func TestIntegerToRomanNegativeNumber(t *testing.T) {
	input := -34
	want := "to raise a (negative number) error"
	got, error := IntegerToRoman(input)
	if error == nil {
		t.Fatalf(`IntegerToRoman("%v") = %v, %v, want %v`, input, got, error, want)
	}
}

func TestIntegerToRomanZero(t *testing.T) {
	input := 0
	want := "to raise a (zero number) error"
	got, error := IntegerToRoman(input)
	if error == nil {
		t.Fatalf(`IntegerToRoman("%v") = %v, %v, want %v`, input, got, error, want)
	}
}
