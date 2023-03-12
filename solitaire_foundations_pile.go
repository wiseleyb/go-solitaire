package main

type SolitaireFoundationsPile struct {
	deck Deck
	suit int
}

func newSolitaireFoundationsPile(suit int) SolitaireFoundationsPile {
	scp := SolitaireFoundationsPile{}
	scp.deck = newDeck()
	scp.suit = suit
	return scp
}
