package main

func main() {

	//var card = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of Diamonds"
	//card := newCard()
	//
	//fmt.Println(card)

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//cards := newDeck()
	//
	////cards.print()
	//
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()

	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting))

	//cards := newDeck()
	//fmt.Println(cards.toString())
	//err := cards.saveToFile("deck.txt")
	//if err != nil {
	//	return
	//}

	// read from file
	//cards := newDeckFrom("deck.txt")
	//cards.print()
	//cards.shuffle()
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
