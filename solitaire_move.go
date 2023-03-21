package main

import (
	"fmt"
)

type SolitaireMove struct {
	sol      *Solitaire
	card     SolitaireCard
	ontoCard SolitaireCard
}

func newSolitaireMove(sol *Solitaire, card SolitaireCard, ontoCard SolitaireCard) SolitaireMove {
	sm := SolitaireMove{}
	sm.sol = sol
	sm.card = card
	sm.ontoCard = ontoCard
	return sm
}

func (sm *SolitaireMove) play() {
	card := sm.pullCard(sm.card)
	disp := newDisplay()
	fmt.Println("Pulled card", disp.card(card))

	switch sm.ontoCard.deckType {
	case "foundation":
		sm.sol.foundations[sm.ontoCard.deckIdx].deck.cards =
			append(sm.sol.foundations[sm.ontoCard.deckIdx].deck.cards, card)
	case "tableau":
		sm.sol.tableaus[sm.ontoCard.deckIdx].shownDeck.cards =
			append(sm.sol.tableaus[sm.ontoCard.deckIdx].shownDeck.cards, card)
	}
}

// TODO: remove card from relevant stack
func (sm *SolitaireMove) pullCard(sc SolitaireCard) Card {
	switch sc.deckType {
	case "foundation":
		card := sm.sol.foundations[sc.deckIdx].deck.last()
		return card
	case "stock":
		card := sm.sol.stock.current()
		return card
	case "tableau":
		// TODO: show next hidden card after pull
		// TODO: handle moving stacks
		card := sm.sol.tableaus[sc.deckIdx].shownDeck.cards[sc.cardIdx]
		return card
	}
	// TODO: should raise error
	card := nullCard()
	return card
}
