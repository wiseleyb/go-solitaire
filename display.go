package main

import (
	"fmt"
	"strings"
)

type Display struct {
	suitNames      [4]string
	suitNamesShort [4]string
	cardNames      [13]string
	cardNamesShort [13]string
}

func newDisplay() Display {
	d := Display{
		suitNames:      [4]string{"Club", "Diamond", "Heart", "Spade"},
		suitNamesShort: [4]string{"♧", "♢", "♡", "♤"},
		cardNames:      [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"},
		cardNamesShort: [13]string{"02", "03", "04", "05", "06", "07", "08", "09", "10", "JJ", "QQ", "KK", "AA"},
	}
	return d
}

func (d Display) deck(deck Deck) string {
	return d.cards(deck.cards)
}
func (d Display) cardLong(c Card) string {
	return d.cardNames[c.value] + " " + d.suitNames[c.suit]
}
func (d Display) card(c Card) string {
	return d.cardNamesShort[c.value] + d.suitNamesShort[c.suit]
}

func (d Display) cards(cards []Card) string {
	var s []string
	for cardIdx := range cards {
		s = append(s, d.card(cards[cardIdx]))
	}
	return strings.Join(s, " ")
}

func (d Display) suitsBorderP() { fmt.Println(d.suitsBorder()) }
func (d Display) suitsBorder() string {
	var s []string
	for x := 0; x < 8; x++ {
		for sidx := 0; sidx < 4; sidx++ {
			s = append(s, d.suitNamesShort[sidx])
		}
	}
	return strings.Join(s, "")
}
