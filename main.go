package main

import "strconv"

func makeRaindropSound(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result += strconv.Itoa(number)
	}

	return result
}

func main() {
	testCases := [4]int{28, 30, 34, 40}

	for i := range len(testCases) {
		println(makeRaindropSound(testCases[i]))
	}

}
