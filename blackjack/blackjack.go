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
	var response string
	var dealerCardValue int = ParseCard(dealerCard)
	var myCardSum int = ParseCard(card1) + ParseCard(card2)

	switch {
	case ParseCard(card1) == 11 && ParseCard(card2) == 11:
		response = "P"
	case myCardSum == 21 && (dealerCardValue != 11 && dealerCardValue != 10):
		response = "W"
	case (myCardSum == 21) && (dealerCardValue == 11 || dealerCardValue == 10):
		response = "S"
	case (myCardSum >= 17 && myCardSum <= 20):
		response = "S"
	case (myCardSum >= 12 && myCardSum <= 16) && dealerCardValue < 7:
		response = "S"
	case (myCardSum >= 12 && myCardSum <= 16) && dealerCardValue >= 7:
		response = "H"
	case myCardSum <= 11:
		response = "H"
	default:
		response = "H"
	}
	return response
}