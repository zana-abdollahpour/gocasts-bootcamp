package main

import "unicode"

func checkIsogram(str string) bool {
	letters := make(map[rune]bool)

	for _, char := range str {
		char = unicode.ToLower(char)

		if !unicode.IsLetter(char) {
			continue
		}

		if letters[char] {
			return false
		}

		letters[char] = true
	}

	return true
}

func main() {
	testcases := [5]string{"lumberjacks", "background", "downstream", "six-year-old", "lululu"}

	for i := range len(testcases) {
		println(checkIsogram(testcases[i]))
	}
}
