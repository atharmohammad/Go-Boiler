package main

func main() {
	cards := newDeck()
	cards.shuffle()
	// left, right := deal(cards, 32)
	// left.print()
	// fmt.Println("------------------------------------------------")
	// right.print()
	cards.save("file")
	cards = read("file")
	cards.print()
}

func newCard() string {
	return "I am your new card"
}
