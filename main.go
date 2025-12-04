package main

import "strings"

func checkValidityViaLuhn(word string) bool {
	num := strings.ReplaceAll(word, " ", "")

	if len(num) <= 1 {
		return false
	}

	lastIndex := len(num) - 1
	sum := 0

	for i, val := range num {
		// Check if character is a digit
		if val < '0' || val > '9' {
			return false
		}

		// Convert rune to actual digit value
		digit := int(val - '0')

		// Double every second digit from the right
		if (lastIndex-i)%2 != 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}

func main() {
	testcases := [4]string{"4539 3195 0343 6467", "066 123 478"}

	for i := range len(testcases) {
		println(checkValidityViaLuhn(testcases[i]))
	}
}
