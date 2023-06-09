package main

import (
	"fmt"
	"strconv"
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
		cardNames:      [13]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"},
		cardNamesShort: [13]string{"AA", "02", "03", "04", "05", "06", "07", "08", "09", "10", "JJ", "QQ", "KK"},
	}
	return d
}

func (d Display) deck(deck Deck) string {
	return d.cards(deck.cards)
}
func (d Display) cardLong(c Card) string {
	if c.value < 0 && c.suit < 0 {
		return "NULL"
	} else {
		return d.cardNames[c.value] + " " + d.suitNames[c.suit]
	}
}
func (d Display) card(c Card) string {
	if c.value < 0 && c.suit < 0 {
		return "NULL"
	} else {
		return d.cardNamesShort[c.value] + d.suitNamesShort[c.suit]
	}
}

func (d Display) solitaire_card(sc SolitaireCard) string {
	return d.card(sc.card)
}

func (d Display) cards(cards []Card) string {
	var s []string
	for cardIdx := range cards {
		s = append(s, d.card(cards[cardIdx]))
	}
	return strings.Join(s, " ")
}

func (d Display) solitaire_cards(solitaire_cards []SolitaireCard) string {
	// TODO: There's got to be a better way to do this
	var cards []Card
	for scIdx := range solitaire_cards {
		sc := solitaire_cards[scIdx]
		cards = append(cards, sc.card)
	}
	return d.cards(cards)
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

func (d Display) move(sp SolitaireMove) string {
	var s []string
	s = append(s, d.solitaire_card(sp.card))
	s = append(s, sp.card.deckType)
	s = append(s, strconv.Itoa(sp.card.deckIdx))
	s = append(s, "->")
	s = append(s, d.solitaire_card(sp.ontoCard))
	s = append(s, sp.ontoCard.deckType)
	s = append(s, strconv.Itoa(sp.ontoCard.deckIdx))
	return strings.Join(s, " ")
}
