package main

import (
	"fmt"
)

type Deck struct {
	cards []Card // 52 shuffled cards
}

func newDeck() Deck {
	return Deck{cards: []Card{}}
}

func new52Deck() Deck {
	d := newDeck()
	for suit := 0; suit <= 3; suit++ {
		for val := 0; val <= 12; val++ {
			d.cards = append(d.cards, newCard(suit, val))
		}
	}
	return d
}

func new52DeckShuffled() Deck {
	d := new52Deck()
	d.shuffle()
	return d
}

func (d *Deck) shuffle() {
	r := newRand()
	cardLen := len(d.cards)
	for x := 0; x < 100; x++ {
		fromIdx := r.rnd(cardLen)
		toIdx := r.rnd(cardLen)
		d.cards[fromIdx], d.cards[toIdx] = d.cards[toIdx], d.cards[fromIdx]
	}
}

func (d *Deck) draw() Card {
	return d.rPop()
}

func (d *Deck) deal(card Card) {
	d.rPush(card)
}

// Push a card before array[0]
func (d *Deck) lPush(card Card) {
	d.reverse()
	d.cards = append(d.cards, card)
	d.reverse()
}

// Push a card after array.last
func (d *Deck) rPush(card Card) {
	d.cards = append(d.cards, card)
}

// removes and returns array[0]
func (d *Deck) lPop() Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

// removes and returns array.last
func (d *Deck) rPop() Card {
	d.reverse()
	card := d.cards[0]
	d.cards = d.cards[1:]
	d.reverse()
	return card
}

func (d Deck) first() Card {
	return d.cards[0]
}

func (d Deck) last() Card {
	return d.cards[len(d.cards)-1]
}

func (d *Deck) reverse() {
	for i, j := 0, len(d.cards)-1; i < j; i, j = i+1, j-1 {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) rotateLastToFirst() {
	card := d.rPop()
	d.lPush(card)
}

func testDeck() {
	deck := new52Deck()
	disp := newDisplay()
	tcnt := 2
	fmt.Println("Original Deck", disp.cards(deck.cards))

	// lPop
	fmt.Println("")
	fmt.Println("test lPop")
	for x := 0; x < 3; x++ {
		fmt.Println("lPop", disp.card(deck.lPop()))
		fmt.Println("Deck", disp.cards(deck.cards))
	}

	// rPop
	fmt.Println("")
	fmt.Println("test rPop")
	for x := 0; x < tcnt; x++ {
		fmt.Println("rPop", disp.card(deck.rPop()))
		fmt.Println("Deck", disp.cards(deck.cards))
	}

	// lPush / rPop
	fmt.Println("")
	fmt.Println("test lPush/rPop")
	for x := 0; x < tcnt; x++ {
		card := deck.rPop()
		fmt.Println("rPop", disp.card(card))
		deck.lPush(card)
		fmt.Println("Deck lPush'd", disp.cards(deck.cards))
	}

	// rPush / lPop
	fmt.Println("")
	fmt.Println("test rPush/lPop")
	for x := 0; x < tcnt; x++ {
		card := deck.lPop()
		fmt.Println("lPop", disp.card(card))
		deck.rPush(card)
		fmt.Println("Deck rPush'd", disp.cards(deck.cards))
	}

	// first / last
	fmt.Println("")
	fmt.Println("test first/last")
	fmt.Println("Deck", disp.cards(deck.cards))
	fmt.Println("first", disp.card(deck.first()))

	// draw
	fmt.Println("")
	fmt.Println("test draw")
	fmt.Println("Deck", disp.cards(deck.cards))
	for x := 0; x < tcnt; x++ {
		card := deck.draw()
		fmt.Println("draw", disp.card(card))
	}
	fmt.Println("Deck", disp.cards(deck.cards))

	// newDeck()
	fmt.Println("")
	fmt.Println("test newDeck()/draw/deal")
	d2 := newDeck()
	fmt.Println("new deck.cards", d2.cards)
	for x := 2; x < 6; x++ {
		card := newCard(1, x)
		fmt.Println("card", disp.card(card))
		d2.lPush(card)
	}
	fmt.Println("lpush'd deck.cards", disp.cards(d2.cards))
}
