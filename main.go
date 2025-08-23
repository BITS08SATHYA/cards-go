package main

import "fmt"

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

	cards := newDeck()
	fmt.Println(cards.toString())

}

func newCard() string {
	return "Five of Diamonds"
}
