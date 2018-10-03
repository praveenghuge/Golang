package main

func main() {
	// var card string = "Ace of Strings"
	cards := newDeck()
	cards.print()
	hand, remainingDeck := deal(cards, 2)
	hand.print()
	remainingDeck.print()

}
func newCard() string {
	return "Ace of Strings"
}
