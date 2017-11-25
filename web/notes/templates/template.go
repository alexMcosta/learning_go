package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Parse a single template if you want to parse more then one
	//Use template.ParseGlob("location/*")
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
