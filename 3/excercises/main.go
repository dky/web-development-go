package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@example.com\">"+
		"support@usegolang.com</a>.")
}

func faqs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>This is the FAQs page, please view"+
	"the questions presented below and let us know via email if you have additional</h2>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(404)
	fmt.Fprint(w, "<h1>We could not find the page you were "+
		"looking for :(</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faqs", faqs)
	//http.ListenAndServe(":3000", nil)
	//Router is passed in here instead of nil.
	http.ListenAndServe(":3000", r)
}
