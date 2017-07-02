package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template


func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeTemplate.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email " +
	"to <a href=\"mailto:support@usegolang.com\">" +
	"support@usegolang.com</a>.")
}
