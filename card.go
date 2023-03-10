package main

type Card struct {
	value int // 0-12 (2,3,4,5,6,7,8,9,10,J,Q,K,A)
	suit  int // 0-3 (club diamon heart spade)
}

func newCard(suit, value int) Card {
	c := Card{
		suit:  suit,
		value: value,
	}
	return c
}
