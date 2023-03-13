package main

type SolitairePlay struct {
	sol           *Solitaire
	possibleMoves []SolitaireMove
}

func newSolitairePlay(sol *Solitaire) SolitairePlay {
	sp := SolitairePlay{}
	sp.sol = sol
	return sp
}

func (sp SolitairePlay) play() {
	sp.sol.display()
	sp.possibleMoves = []SolitaireMove{}
}

// finds cards that could be played
func (sp SolitairePlay) playableCards() []SolitaireCard {
	//disp := newDisplay()
	var res []SolitaireCard
	// check stock
	stock_size := len(sp.sol.stock.deck.cards)
	if stock_size > 0 {
		res = append(res, newSolitaireCard(sp.sol, sp.sol.stock.current(), "stock", -1, -1))
	}
	// check foundations
	for fIdx := range sp.sol.foundations {
		f := sp.sol.foundations[fIdx]
		if len(f.deck.cards) > 0 {
			//fmt.Println("foundation", disp.card(f.deck.last()))
			res = append(res, newSolitaireCard(sp.sol, f.deck.last(), "foundation", fIdx, -1))
		}
	}
	// check tableaus
	for tabIdx := range sp.sol.tableaus {
		tab := sp.sol.tableaus[tabIdx]
		for cardIdx := range tab.shownDeck.cards {
			card := tab.shownDeck.cards[cardIdx]
			res = append(res, newSolitaireCard(sp.sol, card, "tableau", tabIdx, cardIdx))
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
