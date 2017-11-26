package main

import (
	"fmt"
	"net/http"

	//This is what happens when you keeps all your notes and tuts in a centralized directory
	"github.com/alexMcosta/learning_go/web/book_tuts/usegolang.com/chap_6/lenslocked.com/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, e *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is the FAQ page!</h1>")
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>No Page exists</h1>")
}

func main() {

	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	var h http.Handler = http.HandlerFunc(catchAll)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
