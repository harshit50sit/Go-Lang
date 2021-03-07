package main

import "fmt"

var deckSize int

func main() {
	//var card string = "Ace of Spades"
	//or
	deckSize = 20
	card := "Ace of Spades"
	card = "five of Diamonds"
	fmt.Println(card)
	fmt.Println(deckSize)
	//function declaration
	newcard := newCard()
	fmt.Println(newcard)
}

func newCard() string {
	return "seven of jack"
}
