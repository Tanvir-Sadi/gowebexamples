package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/book/{title}/page/{page}", MuxRouting)
	http.ListenAndServe(":8080", r)
}

func MuxRouting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "Title: %s\n Page: %s", title, page)
}
