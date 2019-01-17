package main

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("my_cards")
	cards1, cards2 := deal(cards, 5)
	cards2.saveToFile("my_cards")
	cards1.print()
}
