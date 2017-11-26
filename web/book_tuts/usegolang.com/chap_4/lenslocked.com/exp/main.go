package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name    string
		Age     int
		Sisters []string
	}{
		"Alex Costa",
		29,
		[]string{"Ashley", "Natalie", "Hollishia"}}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
