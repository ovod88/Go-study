package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}

	// cards = append(cards, "Six of Spades")

	// cards.print()
	// cards := newDeck()
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	// X := xa(2)
	// X.print()

	// var y xa = 3 // 3 is an untyped constant assignable to xa
	// y.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand, remainingDeck := cards.deal(5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
