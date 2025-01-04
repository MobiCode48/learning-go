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
func FirstTurn(card1 string, card2 string, dealerCard string) string {

	sumCards := ParseCard(card1) + ParseCard(card2)

	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	if sumCards == 21 {
		if dealerCard != "ace" && dealerCard != "ten" && dealerCard != "queen" && dealerCard != "king" && dealerCard != "jack" {
			return "W"
		}
		return "S"
	}

	if sumCards >= 17 && sumCards <= 20 {
		return "S"
	}

	if sumCards >= 12 && sumCards <= 16 {
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		return "S"
	}

	return "H"
}
