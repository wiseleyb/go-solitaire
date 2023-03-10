package main

type Display struct {
	suitNames [4]string
	cardNames [13]string
}

func newDisplay() Display {
	d := Display{
		suitNames: [4]string{"Club", "Diamond", "Heart", "Spade"},
		cardNames: [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"},
	}
	return d
}

func (d Display) card(c Card) string {
	return d.cardNames[c.value] + " " + d.suitNames[c.suit]
}
