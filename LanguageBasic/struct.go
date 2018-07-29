package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

//Nested struct: struct inside struct, in this way struct CONTACTINFO
//can be reused in other place.
type person struct {
	firstName string
	lastName string
	contact contactInfo
}

//Member method
func (p *person) updateName (name string){
	(*p).firstName = name;
}

func (p person) print () {
	fmt.Printf("%+v", p)
}


func main() {
	//One way
	hiu := person {
		firstName: "Hiu",
		lastName: "Kwok",
		contact: contactInfo{
			email: "hiuk@gmail.com",
			zipCode: 0000,
		},
	}

	hiu.updateName("Andy")
	hiu.print()


	//The other
	contactInfo := contactInfo{ "kk@gmail.com", 0000}
	hiu_2 := person {"Hiu", "K", contactInfo }
	hiu_2.print()

}
