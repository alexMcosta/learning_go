package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct{
	Title string
	Body []byte
}

// Save pages
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load pages
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}