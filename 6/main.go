package main

import (
	"fmt"
	"net/http"
	"./views"

	"github.com/gorilla/mux"
)

var homeView *views.View var contactView *views.View


func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println(homeView.Layout)
	//homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	homeView.Render(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	contactView.Render(w, nil)
}

