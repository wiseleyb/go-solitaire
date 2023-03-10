package main

import (
	"fmt"
)

type Deck struct {
	cards []Card // 52 shuffled cards
	deck  []Card // current deck (decreases as cards are drawn)
	drawn []Card // cards that have been drawn off the deck
}

func newDeck() Deck {
	d := Deck{
		cards: []Card{},
		deck:  []Card{},
		drawn: []Card{},
	}
	//d.cards = make([]Card, 52)
	for suit := 0; suit <= 3; suit++ {
		for val := 0; val <= 12; val++ {
			d.cards = append(d.cards, newCard(suit, val))
			//d.cards[(suit*13)+val] = newCard(suit, val)
		}
	}
	d.shuffle()
	d.deck = d.cards[0:len(d.cards)]

	return d
}

func (d *Deck) drawCard() Card {
	card := d.deck[0]
	d.deck = d.deck[1:]
	d.drawn = append(d.drawn, card)
	return card
}

func (d *Deck) shuffle() {
	fmt.Println("shuffle")
	r := newRand()
	cardLen := len(d.cards)
	for x := 0; x < 100; x++ {
		fromIdx := r.rnd(cardLen)
		toIdx := r.rnd(cardLen)
		fromCard := d.cards[fromIdx]
		toCard := d.cards[toIdx]
		d.cards[toIdx] = fromCard
		d.cards[fromIdx] = toCard
	}
}
