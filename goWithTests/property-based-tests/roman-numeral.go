package main

import (
	// "fmt"
	"strings"
)

type RomanNumeral struct {
	Value int
	Rep string
}

var allRomanNumerals = []RomanNumeral {
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(number int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for number >= numeral.Value{
			result.WriteString(numeral.Rep)
			number -= numeral.Value
		}
	}
	return result.String()

}