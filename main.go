package main

func main() {

	//var card = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of Diamonds"
	//card := newCard()
	//
	//fmt.Println(card)

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
