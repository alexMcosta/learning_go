package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

//To pass a function through a template you need FuncMap
//It Is a composite literal
//It is basically a map for functions
var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	//This is called chaining
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
