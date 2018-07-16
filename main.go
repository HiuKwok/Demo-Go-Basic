//Only name it with main for package would be seat on main executable.
package main

import "fmt"

func main () {
	//No need to specify type String when feeding it with string data.
	//card := "Ace of Spades"
	//:= only work for init, re declare dun work on this way.
	//card = "Five of Diamonds"
	//card := newCard()

	//Turn card into slice
	cards := []string {
			"Ace of Diamonds",
			newCard()}
	//Append
	cards = append (cards, "Six of Spades")

	//Print to String
	//fmt.Println(cards)

	//Iterate
	for i, card := range cards {
		fmt.Println(i, card)
	}


}


func newCard() string {
	return "Five of Diamonds"
}