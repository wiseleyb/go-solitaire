package main

type Card struct {
	value     int // 0-12 (A,2,3,4,5,6,7,8,9,10,J,Q,K)
	suit      int // 0-3 (club diamond heart spade)
	suitNames []string
}

func newCard(suit, value int) Card {
	c := Card{
		suit:  suit,
		value: value,
	}
	c.suitNames = []string{"club", "diamond", "heart", "space"}
	return c
}

func nullCard() Card {
	return Card{suit: -1, value: -1}
}

func (c Card) isAce() bool {
	return c.value == 0
}

func (c Card) isKing() bool {
	return c.value == 12
}

func (c Card) isRed() bool {
	return c.color() == "red"
}

func (c Card) isBlack() bool {
	return c.color() == "black"
}

func (c Card) color() string {
	if c.suit == 1 || c.suit == 2 {
		return "red"
	} else {
		return "black"
	}
}

func (c Card) suitName() string {
	return c.suitNames[c.suit]
}

func (c Card) notEqualTo(card Card) bool {
	return !(c.value == card.value && c.suit == card.suit)
}
