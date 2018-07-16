//Only name it with main for package would be seat on main executable.
package main

func main () {


	//Turn card into slice
	cards := newDeck()
	//Append
	cards = append (cards, "Six of Spades")

	cards.print()


}

