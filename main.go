package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>we could not find the page you were looking for :(</h1><p>please email us if you keep being send to an invalid page.</p>")
	}
}
func main() {
	fmt.Println("server running on port 3000")

	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/contact", handlerFunc)
	http.ListenAndServe(":3000", r)
}
