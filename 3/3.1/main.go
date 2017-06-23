package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<h1>Hello, welcome to my awesome site!</h1>")
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
		//fmt.Println("Does this even work?")
	} else  if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@example.com\">"+
		"support@usegolang.com</a>.")
	} else {
		//statusNotFound = constant for 404
		//w.WriteHeader(http.StatusNotFound)
		w.WriteHeader(404)
		fmt.Fprint(w, "<h1>We could not find the page you were "+
		"looking for :(</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
