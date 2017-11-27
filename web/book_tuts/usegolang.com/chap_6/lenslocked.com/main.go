package main

import (
	"net/http"

	//This is what happens when you keeps all your notes and tuts in a centralized directory
	"github.com/alexMcosta/learning_go/web/book_tuts/usegolang.com/chap_6/lenslocked.com/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
var catchAllView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, e *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(catchAllView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	catchAllView = views.NewView("bootstrap", "views/catchall.gohtml")

	var h http.Handler = http.HandlerFunc(catchAll)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
