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

	//cards := newDeck()
	//cards.shuffleq()
	//cards.print()
	//
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}

}

func newCard() string {
	return "Five of Diamonds"
}
