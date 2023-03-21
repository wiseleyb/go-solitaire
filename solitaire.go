package main

import (
	"fmt"
)

type Solitaire struct {
	deck          Deck
	stock         SolitaireStock
	suits         [4]string
	foundations   [4]SolitaireFoundationsPile // aces stacks
	tableaus      [7]SolitaireTableauPile
	playableCards []SolitaireCard
	possibleMoves []SolitaireMove
	playedMoves   []SolitaireMove
	moves         int
}

func newSolitaire() Solitaire {
	sol := Solitaire{}
	sol.restart()
	return sol
}

func (sol *Solitaire) restart() {
	sol.deck = new52DeckShuffled()
	sol.suits = [4]string{"Club", "Diamond", "Heart", "Space"}
	for suitIdx := range sol.suits {
		sol.foundations[suitIdx] = newSolitaireFoundationsPile(suitIdx)
	}
	for tabIdx := range sol.tableaus {
		for x := 0; x < tabIdx; x++ {
			card := sol.deck.draw()
			sol.tableaus[tabIdx].hiddenDeck.lPush(card)
		}
		card := sol.deck.draw()
		sol.tableaus[tabIdx].shownDeck.lPush(card)
	}
	sol.stock = newSolitaireStock(sol.deck)
	sol.moves = 0
}

func (sol *Solitaire) move() {
	sp := newSolitairePlay(sol)
	sp.move()
}

func (sol Solitaire) stockCard() Card {
	return sol.stock.deck.last()
}

// returns a foundation deck for a given suit
func (sol Solitaire) foundationDeckBySuit(suit int) Deck {
	for fdIdx := range sol.foundations {
		if sol.foundations[fdIdx].suit == suit {
			return sol.foundations[fdIdx].deck
		}
	}
	// TODO: how to return nil/error
	return newDeck()
}

func (sol *Solitaire) display() {
	disp := newDisplay()
	fmt.Println("")
	fmt.Println("")
	disp.suitsBorderP()
	fmt.Println("Solitaire Board", "moves", sol.moves)
	fmt.Println("")
	fmt.Println("Foundations")
	for suitIdx := range sol.suits {
		fmt.Println("Ace", sol.suits[suitIdx], disp.deck(sol.foundations[suitIdx].deck))
	}

	fmt.Println("")
	fmt.Println("Tableaus")
	for idx := 0; idx < len(sol.tableaus); idx++ {
		fmt.Println(idx,
			"Shown",
			disp.deck(sol.tableaus[idx].shownDeck),
			"Hidden",
			disp.deck(sol.tableaus[idx].hiddenDeck))
	}

	fmt.Println("")
	fmt.Println("Stock", 
              disp.card(sol.stock.current()), 
              "Stock Deck", 
              disp.deck(sol.stock.deck))

	fmt.Println("")
	fmt.Println("")
}
