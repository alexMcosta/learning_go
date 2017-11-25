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

	//Map we are sending to the template
	favorites := map[string]string{
		"Ice Cream": "Cookies and Cream",
		"Cookie":    "Oreo",
		"Animal":    "Cat",
		"OS":        "Ubuntu",
		"Programming Language": "Go"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", favorites)
	if err != nil {
		log.Fatalln(err)
	}
}
