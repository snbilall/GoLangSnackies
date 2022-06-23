package main

import "gosnackies/cards"

func main() {
	cards := cards.NewDeck()
	cards.Shuffle()
	cards.Print()
}
