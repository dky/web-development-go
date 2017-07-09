package main

import (
	"fmt"
	"net/http"
	"./views"

	"github.com/gorilla/mux"
)

var homeView *views.View 
var contactView *views.View
var faqsView *views.View
var signupView *views.View

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqsView = views.NewView("bootstrap", "views/faqs.gohtml")
	signupView = views.NewView("bootstrap", "views/signup.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faqs", faqs)
	r.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", r)
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	signupView.Render(w, nil)
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

func faqs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	faqsView.Render(w, nil)
}
