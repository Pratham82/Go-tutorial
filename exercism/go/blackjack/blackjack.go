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
	card1Parsed := ParseCard(card1)
	card2Parsed := ParseCard(card2)
	dealerCardParsed := ParseCard(dealerCard)

	switch sum := card1Parsed + card2Parsed; {
	case card1Parsed == 11 && card2Parsed == 11:
		return "P"
	case card1Parsed+card2Parsed == 21 && (dealerCardParsed == 10 || dealerCardParsed == 11):
		return "S"
	case card1Parsed+card2Parsed == 21:
		return "W"
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16 && dealerCardParsed >= 7:
		return "H"
	case sum >= 12 && sum <= 16:
		return "S"
	case sum <= 12:
		return "H"
	default:
		return "H"
	}

}
