package main

type SolitaireStock struct {
	deck Deck
}

func newSolitaireStock(deck Deck) SolitaireStock {
	ss := SolitaireStock{}
	ss.deck = deck //new52DeckShuffled()
	return ss
}

func (sp SolitaireStock) current() Card {
	return sp.deck.last()
}

// moves last card to first card
// returns last card
func (sp *SolitaireStock) next() Card {
	sp.deck.rotateLastToFirst()
	return sp.current()
}
