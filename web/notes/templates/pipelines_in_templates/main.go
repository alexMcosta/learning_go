package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funcmap).ParseFiles("tpl.gohtml"))
}

func addOne(x int) int {
	return x + 1
}

func subTwo(x int) int {
	return x - 2
}

func dbl(x int) int {
	return x + x
}

var funcmap = template.FuncMap{
	"addOne": addOne,
	"subTwo": subTwo,
	"dbl":    dbl,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
