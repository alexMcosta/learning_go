package main

import "fmt"

//OOP in Go is the use of structs which are not
// Objects but a new type
type coworker struct {
	firstname string
	lastame   string
	age       int
	developer bool
}

//Methods in in go are like so
func (p coworker) fullname() string {
	return p.firstname + " " + p.lastame
}

func main() {
	corey := coworker{"Corey", "Wellington", 28, true}
	nicholas := coworker{"Nick", "Parker", 37, true}
	sam := coworker{"Sam", "Matthews", 34, false}
	kyle := coworker{"Kyle", "Fletcher", 30, false}
	chris := coworker{"Chris", "Packer", 29, false}
	matt := coworker{"Matt", "Stint", 29, false}
	packer := coworker{chris.lastame, nicholas.lastame, corey.age, matt.developer}

	fmt.Println(corey.firstname, nicholas.lastame, sam.age, kyle.developer, packer, corey.fullname())
}
