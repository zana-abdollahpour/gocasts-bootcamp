package main

func calculateHammingDistance(seq1, seq2 string) int {
	if len(seq1) != len(seq2) {
		return 0
	}

	result := 0

	for idx := range len(seq1) {
		if seq1[idx] != seq2[idx] {
			result++
		}

	}

	return result
}

func main() {
	strand1 := "GAGCCTACTAACGGGAT"
	strand2 := "CATCGTAATGACGGCCT"

	println(calculateHammingDistance(strand1, strand2))
}
