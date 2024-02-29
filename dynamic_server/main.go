package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello From Http Server")
	http.HandleFunc("/", DynamicServer)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
func DynamicServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From Dynamic HTTP Server")
}
