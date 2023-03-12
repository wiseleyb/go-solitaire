package main

import (
	"fmt"
)

type Solitaire struct {
	stock       Deck
	suits       [4]string
	foundations [4]SolitaireFoundationsPile // aces stacks
	stockIdx    int
	tableaus    [7]SolitaireTableauPile
	moves       int
}

func newSolitaire() Solitaire {
	sol := Solitaire{}
	sol.restart()
	return sol
}

func (sol *Solitaire) play() {
	sp := newSolitairePlay(sol)
	sp.play()
}

func (sol Solitaire) stockCard() Card {
	return sol.stock.cards[sol.stockIdx]
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

func (sol *Solitaire) restart() {
	sol.stock = new52DeckShuffled()
	sol.suits = [4]string{"Club", "Diamond", "Heart", "Space"}
	for suitIdx := range sol.suits {
		sol.foundations[suitIdx] = newSolitaireFoundationsPile(suitIdx)
	}
	sol.moves = 0
	fmt.Println("stock", sol.stock)
	for tabIdx := range sol.tableaus {
		for x := 0; x < tabIdx; x++ {
			card := sol.stock.draw()
			sol.tableaus[tabIdx].hiddenDeck.lPush(card)
		}
		card := sol.stock.draw()
		sol.tableaus[tabIdx].shownDeck.lPush(card)
	}
}

func (sol *Solitaire) display() {
	disp := newDisplay()
	fmt.Println("")
	fmt.Println("")
	disp.suitsBorderP()
	fmt.Println("Solitaire Board")
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
	fmt.Println("Moves", sol.moves)

	fmt.Println("")
	fmt.Println("Stock", disp.card(sol.stock.cards[sol.stockIdx]), "Stock Deck", disp.deck(sol.stock))

	fmt.Println("")
	sp := newSolitairePlay(sol)
	pcs := sp.playableCards()
	fmt.Println("Playable Cards", disp.cards(pcs))

	fmt.Println("")
	for cardIdx := range pcs {
		card := pcs[cardIdx]
		fmt.Println(disp.card(card))
		moves := sp.playCard(card)
		for moveIdx := range moves {
			fmt.Println("   ", disp.move(moves[moveIdx]))
		}
	}

	fmt.Println("")
	fmt.Println("")
}
