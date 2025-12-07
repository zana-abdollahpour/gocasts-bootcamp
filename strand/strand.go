package strand

var complementaryStrands = map[rune]rune{
	'A': 'U',
	'T': 'A',
	'C': 'G',
	'G': 'C',
}

func ToRNA(dna string) string {
	result := ""

	for _, nucleotide := range dna {
		result += string(complementaryStrands[nucleotide])
	}

	return result
}
