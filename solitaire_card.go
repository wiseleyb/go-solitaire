package main

// TODO: this would be cleaner by passing references... need to learn how to optional args in Go
type SolitaireCard struct {
	sol      *Solitaire
	card     Card
	deckType string
	deckIdx  int
	cardIdx  int
}

func newSolitaireCard(sol *Solitaire, card Card, deckType string, deckIdx int, cardIdx int) SolitaireCard {
	sc := SolitaireCard{}
	sc.sol = sol
	sc.card = card
	sc.deckType = deckType
	sc.deckIdx = deckIdx
	sc.cardIdx = cardIdx
	return sc
}

func nullSolitaireCard(sol *Solitaire, deckType string, deckIdx int) SolitaireCard {
	sc := SolitaireCard{}
	sc.sol = sol
	sc.card = nullCard()
	sc.deckType = deckType
	sc.deckIdx = deckIdx
	sc.cardIdx = -1
	return sc
}
