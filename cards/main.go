package main

func main() {
	cards := newDeck()
	cards.shuffle()
	// cards.saveToFile("my_cards")
	cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("remainingCards")
	// remainingCards.print()

}
