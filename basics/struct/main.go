package main

import "fmt"

//OOP in Go is the use of structs which are not
// Objects but a new type
type Coworker struct {
	firstname string
	lastname  string
	age       int
	developer bool
}

//Methods in in go are like so
func (p Coworker) fullname() string {
	return p.firstname + " " + p.lastname
}

//Inheritance
type CurrentCoworker struct {
	Coworker
	Current bool
}

func main() {
	corey := Coworker{"Corey", "Wellington", 28, true}
	nicholas := Coworker{"Nick", "Parker", 37, true}
	sam := Coworker{"Sam", "Matthews", 34, false}
	kyle := Coworker{"Kyle", "Fletcher", 30, false}
	chris := Coworker{"Chris", "Packer", 29, false}
	matt := Coworker{"Matt", "Stint", 29, false}
	packer := Coworker{chris.lastname, nicholas.lastname, corey.age, matt.developer}

	//Creating a struct that has inheritance
	eb := CurrentCoworker{
		Coworker: Coworker{
			firstname: "Elizabeth",
			lastname:  "Bye",
			age:       50,
			developer: false,
		},
		Current: false,
	}

	fmt.Println(corey.firstname, nicholas.lastname, sam.age, kyle.developer, packer, corey.fullname(), eb.firstname, eb.Current)
}
