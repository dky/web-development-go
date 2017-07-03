package main

import (
	"net/http"
	"./views"

	"github.com/gorilla/mux"
)

/*
var homeTemplate *template.Template
var contactTemplate *template.Template
*/

var homeView *views.View
var contactView *views.View


func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//homeTemplate.Execute(w, nil)
	homeView.Template.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//contactTemplate.Execute(w, nil)
	contactView.Template.Execute(w, nil)
}
