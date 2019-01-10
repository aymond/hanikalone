package main

import "fmt"
import "gocards"

func main() {
	deck, _ := gocards.NewDeck("Test")
	deck.Debug()

	fmt.Println("test")

	player1, _ := gocards.NewDeck("Player1")
	player2, _ := gocards.NewDeck("Player2")
	discard, _ := gocards.NewDeck("Discard")

	player1.Debug()
	player2.Debug()
	discard.Debug()
	deck.Cards = gocards.GenerateDeck("hanikalone")
	deck.Debug()
	deck.Shuffle()
	deck.Debug()
}
