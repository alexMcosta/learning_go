package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	//Slice we are sending to the template
	friends := []string{"Tessa", "Corey", "Nicholas", "Squeek", "Kyle", "Jake"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", friends)
	if err != nil {
		log.Fatalln(err)
	}
}
