package main

func calcDifferenceOfSquares(num int) int {
	sum := 0
	squareOfTheSum := 0
	sumOfTheSquares := 0

	for i := range num + 1 {
		if i == 0 {
			continue
		}

		sum += i
		sumOfTheSquares += i * i
	}

	squareOfTheSum = sum * sum

	return squareOfTheSum - sumOfTheSquares
}

func main() {
	testcases := [4]int{10, 4, 14, 40}

	for i := range len(testcases) {
		println(calcDifferenceOfSquares(testcases[i]))
	}
}
