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
	sol.display()
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
	// play turn
	sol.moves++
	disp := newDisplay()
	disp.suitsBorderP()
	fmt.Println("")
	fmt.Println("PLAY TURN")
	sol.findMoves()
	sol.displayPossibleMoves()
	sol.playMove()
	sol.display()
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

// resets possibleMoves
// populates possibleMoves given current board
func (sol *Solitaire) findMoves() {
	sp := newSolitairePlay(sol)
	pcs := sp.playableCards()
	sol.playableCards = pcs
	sol.possibleMoves = []SolitaireMove{}
	for cardIdx := range pcs {
		card := pcs[cardIdx]
		moves := sp.findMoves(card)
		for moveIdx := range moves {
			sol.possibleMoves = append(sol.possibleMoves, moves[moveIdx])
		}
	}
}

// plays a possibleMove
func (sol *Solitaire) playMove() {
	if len(sol.possibleMoves) > 0 {
		move := sol.possibleMoves[0]
		move.play()
		sol.playedMoves = append(sol.playedMoves, move)
	}
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
	fmt.Println("Stock", disp.card(sol.stock.current()), "Stock Deck", disp.deck(sol.stock.deck))

	fmt.Println("")
	fmt.Println("")
}

func (sol *Solitaire) displayPossibleMoves() {
	disp := newDisplay()
	fmt.Println("")
	fmt.Println("Playable Cards", disp.solitaire_cards(sol.playableCards))
	for moveIdx := range sol.possibleMoves {
		move := sol.possibleMoves[moveIdx]
		fmt.Println(disp.move(move))
	}
}
