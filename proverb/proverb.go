package proverb

// Proverb generates the proverb rhyme from a list of words
func Proverb(rhyme []string) []string {
	// Handle edge case: empty input
	if len(rhyme) == 0 {
		return []string{}
	}

	result := make([]string, 0, len(rhyme))

	// Generate the "For want of..." lines for consecutive pairs
	for i := 0; i < len(rhyme)-1; i++ {
		line := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		result = append(result, line)
	}

	// Add the final line using the first word
	finalLine := "And all for the want of a " + rhyme[0] + "."
	result = append(result, finalLine)

	return result
}
