package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

//Decleration of the struct we are sending
type pet struct {
	Name    string
	Species string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	// Struct we are sending to template
	sausages := pet{
		Name:    "Sausages",
		Species: "Ferret",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sausages)
	if err != nil {
		log.Fatalln(err)
	}
}
