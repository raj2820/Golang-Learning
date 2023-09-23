package main

func main() {
	// var card string = "Ace of Spaces"
	// card := newCard()
	// card = "Five of Diamonds"
	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "six of spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards := newDeck()
	// cards.print()
	// fmt.Println(cards)

		// ******************** multiple return values
	hand, remainingCards := deal(cards, 5)
  
	hand.print()
	remainingCards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
