package main

import (
	"fmt"
	"net/http"

	"./views"
	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func notFound404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-html")
	fmt.Fprint(w, "<h1>FAQ</h1> if you have any question please select category below")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404</h1> sorry your page not found")
}

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/", home)

	http.ListenAndServe(":3000", r)
}
