package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    value := 0
	switch {
        case card=="ace":
        	value = 11
        case card=="king" || card=="queen" || card=="jack" || card=="ten":
        	value = 10
        case card=="nine":
        	value = 9
        case card=="eight":
        	value = 8
        case card=="seven":
        	value=7
        case card=="six":
        	value=6
        case card=="five":
        	value=5
        case card=="four":
        	value=4
        case card=="three":
        	value=3
        case card=="two":
        	value=2
        default:
        	value=0
    }

    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
    dv := ParseCard(dealerCard)

    var decision string
    
    switch {
        case sum==22:
        	decision = "P"
        case sum==21 && dv < 10:
        	decision = "W"
        case sum==21 && dv >= 10:
        	decision = "S"
		case sum>=17:
        	decision = "S"
        case sum>=12 && sum<=16 && dv>=7:
        	decision = "H"
        case sum>=12 && sum<=16:
        	decision = "S"
        case sum<=11:
        	decision = "H"
    }

    return decision
}
