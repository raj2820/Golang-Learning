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


	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards.print()
	// fmt.Println(cards)

	// ******************** multiple return values
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

		// cards := newDeckFromFile("my")
		// cards.print()

		cards := newDeck()
		cards.shuffel()
		cards.print()
	}


// func newCard() string {
// 	return "Five of Diamonds"
// }
