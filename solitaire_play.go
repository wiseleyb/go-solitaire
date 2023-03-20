package main

import "fmt"

type SolitairePlay struct {
	sol           *Solitaire
	possibleMoves []SolitaireMove
}

func newSolitairePlay(sol *Solitaire) SolitairePlay {
	sp := SolitairePlay{}
	sp.sol = sol
	return sp
}

func (sp SolitairePlay) move() {
	// play turn
	sp.sol.moves++
	disp := newDisplay()
	disp.suitsBorderP()
	fmt.Println("")
	fmt.Println("PLAY TURN")
	sp.findAllMoves()
	sp.displayPossibleMoves()
	sp.playMove()
}

// resets possibleMoves
// populates possibleMoves given current board
func (sp SolitairePlay) findAllMoves() {
	pcs := sp.playableCards()
	sp.sol.playableCards = pcs
	sp.sol.possibleMoves = []SolitaireMove{}
	for cardIdx := range pcs {
		card := pcs[cardIdx]
		moves := sp.findMoves(card)
		for moveIdx := range moves {
			sp.sol.possibleMoves = append(sp.sol.possibleMoves, moves[moveIdx])
		}
	}
}

// finds cards that could be played
func (sp SolitairePlay) playableCards() []SolitaireCard {
	//disp := newDisplay()
	var res []SolitaireCard
	// check stock
	stock_size := len(sp.sol.stock.deck.cards)
	if stock_size > 0 {
    nsc := newSolitaireCard(sp.sol, sp.sol.stock.current(), "stock", -1, -1)
		res = append(res, nsc)
	}
	// check foundations
	for fIdx := range sp.sol.foundations {
		f := sp.sol.foundations[fIdx]
		if len(f.deck.cards) > 0 {
			//fmt.Println("foundation", disp.card(f.deck.last()))
      nsc := newSolitaireCard(sp.sol, f.deck.last(), "foundation", fIdx, -1)
			res = append(res, nsc)
		}
	}
	// check tableaus
	for tabIdx := range sp.sol.tableaus {
		tab := sp.sol.tableaus[tabIdx]
		for cardIdx := range tab.shownDeck.cards {
			card := tab.shownDeck.cards[cardIdx]
      nsc := newSolitaireCard(sp.sol, card, "tableau", tabIdx, cardIdx)
			res = append(res, nsc)
		}
	}
	return res
}

// given a card - returns data on where it could be played
func (sp SolitairePlay) findMoves(sc SolitaireCard) []SolitaireMove {
	var res []SolitaireMove
	// check foundations
	for fIdx := range sp.sol.foundations {
		f := sp.sol.foundations[fIdx]
		if sc.card.suit == f.suit {
			//   foundation empty && card is ace && suits match
			if len(f.deck.cards) == 0 && sc.card.isAce() {
				res = append(res, newSolitaireMove(sp.sol, sc, nullSolitaireCard(sp.sol, "foundation", fIdx)))
			}
			//   foundation has card, suits match, and card.val = foundation.card.val + 1
			if len(f.deck.cards) > 0 && f.deck.last().value == sc.card.value-1 {
				ontoSc := newSolitaireCard(sp.sol, f.deck.last(), "foundation", fIdx, -1)
				res = append(res, newSolitaireMove(sp.sol, sc, ontoSc))
			}
		}
	}
	// check tableaus
	for tabIdx := range sp.sol.tableaus {
		tab := sp.sol.tableaus[tabIdx]
		//   showDeck.shownCard.nil && card.king
		if len(tab.shownDeck.cards) == 0 && sc.card.isKing() {
			res = append(res, newSolitaireMove(sp.sol, sc, nullSolitaireCard(sp.sol, "tableau", tabIdx)))
		}
		//   shownDeck.shownCard(val, suit) = card.val + 1 and card.suit.color <> card.color
		for cardIdx := range tab.shownDeck.cards {
			c := tab.shownDeck.cards[cardIdx]
			if c.notEqualTo(sc.card) &&
				sc.card.value == c.value-1 &&
				sc.card.color() != c.color() &&
				!(sc.card.isAce()) {
				ontoSc := newSolitaireCard(sp.sol, c, "tableau", tabIdx, cardIdx)
				res = append(res, newSolitaireMove(sp.sol, sc, ontoSc))
			}
		}
	}
	return res
}

// plays a possibleMove
func (sp SolitairePlay) playMove() {
	if len(sp.sol.possibleMoves) > 0 {
		move := sp.sol.possibleMoves[0]
		move.play()
		sp.sol.playedMoves = append(sp.sol.playedMoves, move)
	}
}

func (sp SolitairePlay) displayPossibleMoves() {
	disp := newDisplay()
	fmt.Println("")
	fmt.Println("Playable Cards", disp.solitaire_cards(sp.sol.playableCards))
	for moveIdx := range sp.sol.possibleMoves {
		move := sp.sol.possibleMoves[moveIdx]
		fmt.Println(disp.move(move))
	}
}
