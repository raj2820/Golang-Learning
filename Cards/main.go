package main

import "fmt"

func main() {
	// var card string = "Ace of Spaces"
	// card := newCard()
	// card = "Five of Diamonds"
	cards := []string{newCard(),newCard()} 
	cards = append(cards, "six of spades")
	for i, card := range cards{
		fmt.Println(i,card)
	}
	
	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}

