package main

import (
	"fmt"
)

func main() {
	fmt.Println("cards")
	c := newCard(1, 2)
	fmt.Println(c)
	d := newDeck()
	fmt.Println(d)
	disp := newDisplay()
	fmt.Println(disp.suitNames[0])
	fmt.Println("Card: ", disp.card(c))
	fmt.Println(len(d.deck))
	fmt.Println("Deck", len(d.deck), d)
	fmt.Println("", "")
	for x := 0; x < 55; x++ {
		fmt.Println("Draw", x)
		fmt.Println(d.drawCard())
		fmt.Println("Cards", d.cards)
		fmt.Println("Deck", d.deck)
		fmt.Println("Drawn", d.drawn)
		fmt.Println("", "")
	}
}
