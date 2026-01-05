package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	dealerValue := ParseCard(dealerCard)
	userValue1, userValue2 := ParseCard(card1), ParseCard(card2)
	userSumValue := userValue1 + userValue2

	switch {
	case userSumValue == 22:
		return "P"
	case userValue1+userValue2 == 21 && dealerValue < 10:
		return "W"
	case userSumValue >= 17 && userSumValue <= 20:
		return "S"
	case userSumValue >= 12 && userSumValue <= 16 && dealerValue < 7:
		return "S"
	case userSumValue >= 12 && userSumValue <= 16 && dealerValue >= 7:
		return "H"
	case userSumValue <= 11:
		return "H"
	default:
		return "S"
	}
}
