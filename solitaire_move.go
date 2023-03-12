package main

import "fmt"

type SolitaireMove struct {
	sol      *Solitaire
	card     Card
	ontoCard Card
	deckType string
	deckIdx  int
}

func newSolitaireMove(sol *Solitaire, deckType string, deckIdx int, card Card, ontoCard Card) SolitaireMove {
	sm := SolitaireMove{}
	sm.sol = sol
	sm.deckType = deckType
	sm.deckIdx = deckIdx
	sm.card = card
	sm.ontoCard = ontoCard
	return sm
}

func (sm SolitaireMove) display() {
	fmt.Println(sm)
}
