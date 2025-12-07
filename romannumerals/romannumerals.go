package romannumerals

import (
	"errors"
	"strings"
)

// ToRomanNumeral converts an Arabic number to a Roman numeral string
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("input must be between 1 and 3000")
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder

	for i := range values {
		for input >= values[i] {
			result.WriteString(numerals[i])
			input -= values[i]
		}
	}

	return result.String(), nil
}
