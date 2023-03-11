package main

type SolitaireFoundationsPile struct {
	deck Deck
	suit string
}

func newSolitaireFoundationsPile(suit string) SolitaireFoundationsPile {
	scp := SolitaireFoundationsPile{}
	scp.deck = newDeck()
	scp.suit = suit
	return scp
}
