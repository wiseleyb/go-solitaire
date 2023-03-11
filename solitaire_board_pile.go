package main

type SolitaireTableauPile struct {
	hiddenDeck Deck
	shownDeck  Deck
}

func newSolitaireTableauPile() SolitaireTableauPile {
	stp := SolitaireTableauPile{}
	stp.hiddenDeck = newDeck()
	stp.shownDeck = newDeck()
	return stp
}
