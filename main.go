//Only name it with main for package would be seat on main executable.
package main

import "fmt"

func main () {

	cards := newDeck()
	cards.shuffle()
	fmt.Println(cards)


}

