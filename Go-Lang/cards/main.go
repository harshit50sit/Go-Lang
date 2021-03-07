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
	//Slices & for loops
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") //add a new element
	fmt.Println(cards)
	//Iterate over slice of cards

	for i, card := range cards {
		fmt.Println(i, card)
	}
	//2nd metode
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func newCard() string {
	return "seven of jack"
}
