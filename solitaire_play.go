package main

import "fmt"

type SolitairePlay struct {
	sol *Solitaire
}

func newSolitairePlay(sol *Solitaire) SolitairePlay {
	sp := SolitairePlay{}
	sp.sol = sol
	return sp
}

func (sp SolitairePlay) play() {
	sp.sol.display()
	sp.playAces()
}

// finds cards that could be played
func (sp SolitairePlay) playableCards() []Card {
	//disp := newDisplay()
	var res []Card
	// check stock
	if len(sp.sol.stock.cards) > 0 {
		//fmt.Println("stock", disp.card(sp.sol.stock.last()))
		res = append(res, sp.sol.stock.last())
	}
	// check foundations
	for fIdx := range sp.sol.foundations {
		f := sp.sol.foundations[fIdx]
		if len(f.deck.cards) > 0 {
			//fmt.Println("foundation", disp.card(f.deck.last()))
			res = append(res, f.deck.last())
		}
	}
	// check tableaus
	for tabIdx := range sp.sol.tableaus {
		tab := sp.sol.tableaus[tabIdx]
		for cardIdx := range tab.shownDeck.cards {
			//fmt.Println("tab", tabIdx, cardIdx, disp.card(tab.shownDeck.cards[cardIdx]))
			res = append(res, tab.shownDeck.cards[cardIdx])
		}
	}
	return res
}

// given a card - returns data on where it could be played
func (sp SolitairePlay) playCard(card Card) []SolitaireMove {
	var res []SolitaireMove
	// check foundations
	for fIdx := range sp.sol.foundations {
		f := sp.sol.foundations[fIdx]
		if card.suit == f.suit {
			//   foundation empty && card is ace && suits match
			if len(f.deck.cards) == 0 && card.isAce() {
				res = append(res, newSolitaireMove(sp.sol, "foundation", fIdx, card, nullCard()))
			}
			//   foundation has card, suits match, and card.val = foundation.card.val + 1
			if len(f.deck.cards) > 0 && f.deck.last().value == card.value-1 {
				res = append(res, newSolitaireMove(sp.sol, "foundation", fIdx, card, f.deck.last()))
			}
		}
	}
	// check tableaus
	for tabIdx := range sp.sol.tableaus {
		tab := sp.sol.tableaus[tabIdx]
		//   showDeck.shownCard.nil && card.king
		if len(tab.shownDeck.cards) == 0 && card.isKing() {
			res = append(res, newSolitaireMove(sp.sol, "tableau", tabIdx, card, nullCard()))
		}
		//   shownDeck.shownCard(val, suit) = card.val + 1 and card.suit.color <> card.color
		for cardIdx := range tab.shownDeck.cards {
			c := tab.shownDeck.cards[cardIdx]
			if c.notEqualTo(card) && card.value == c.value-1 && card.color() != c.color() {
				res = append(res, newSolitaireMove(sp.sol, "tableau", tabIdx, card, c))
			}
		}
	}
	return res
}

// If a ace can be played, it plays it and returns true
// Any aces? Check stock + last up card
// if all foundations have one card - bail
// check displayed stock card for ace
// check tableaus for ace
func (sp *SolitairePlay) playAces() bool {
	// Check if any aces can be played?
	if sp.sol.stockCard().isAce() {
		fmt.Println("stock ace", sp.sol.stockCard())
		return true
	}
	for stpIdx := range sp.sol.tableaus {
		tab := sp.sol.tableaus[stpIdx]
		if tab.shownCard().isAce() {
			fmt.Println("tableau ace", stpIdx, tab.shownCard())
			return true
		}
	}
	return false
	//TODO: better way to check All Foundations
}
