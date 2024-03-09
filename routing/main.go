package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/{title}/page/{page}", MuxRouting)
	bookRouter.HandleFunc("/{title}", ReadBook).Methods("GET")
	bookRouter.HandleFunc("/{title}", CreateBook).Methods("POST")
	bookRouter.HandleFunc("/{title}", UpdateBook).Methods("PUT")
	bookRouter.HandleFunc("/{title}", DeleteBook).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func MuxRouting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "Title: %s\nPage: %s", title, page)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "Title: %s\n", title)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
