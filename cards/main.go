package main

import "fmt"

func main() {
	cards := newDeck()
	// we can add element to existing slice with append function
	cards = append(cards, "Magma Opus")

	fmt.Println("---------------Prismari Midrange---------------")
	cards.print()
	fmt.Println("-----------------------------------")

	hand, remainingDeck := deal(cards, 3)

	fmt.Println("***************Hand***************")
	hand.print()
	fmt.Println("***************Remaining deck***************")
	remainingDeck.print()
	fmt.Println("***************Cards to string***************")
	println(cards.toString())
	fmt.Println("***************Save deck to file***************")
	cards.saveToFile("my_test_deck")
	fmt.Println("***************Read deck from file***************")
	cardsFromFile := newDeckFromFile("deck.txt")
	cardsFromFile.print()
	fmt.Println("***************Shuffle deck***************")
	cardsFromFile.shuffle()
	cardsFromFile.print()
}
