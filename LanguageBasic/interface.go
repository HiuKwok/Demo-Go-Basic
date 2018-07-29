package main

import "fmt"

//Define a interface
type Bot interface {
	greet(name string)
}


//New type to implement such interface
type CantoneseBot  struct {

}
//Define a method which call
func(bot CantoneseBot) greet(name string) {
	fmt.Println("Hello " + name);
}

//Func that would take a interface and call relevant func which defined
//on interface.
func callGreet (bot Bot) {
	bot.greet("Hiu");
}

//Main func that construct a new implementation then pass into above
//which take interface as input. Then in compile-time, compiler would
//determine that: Oh implementation have implement all definition on
//interface, so all good to go and compiler would let it pass.
func main() {
	cbot := CantoneseBot{}
	callGreet(&cbot);
}

