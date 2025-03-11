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
		default :
			return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerFirstCard := ParseCard(card1)
	playerSecondCard := ParseCard(card2)
	playerTotal := playerFirstCard + playerSecondCard
	dealer := ParseCard(dealerCard)

	switch {
		case playerFirstCard == 11 && playerSecondCard == 11 || card1 == "ace" && card2 == "ace":
			return "P"
		case playerTotal == 21 && dealer < 10:
			return "W"
		case playerTotal == 21 && dealer >= 10:
			return "S"
		case playerTotal >= 17 && playerTotal <= 20 :
			return "S"
		case playerTotal >= 12 && playerTotal <= 16 && dealer >= 7:
			return "H"
		case playerTotal >= 0 && playerTotal <= 11:
			return "H"
		case playerTotal == 21 && dealer <= 20 :
			return "W"
		default :
			return "S"
	}
	
}
