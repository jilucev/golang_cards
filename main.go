package main

func main() {
	// cards := newDeckFromFile("my_cards")
	// fmt.Println(cards)

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
