package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/",HomePage)
	http.ListenAndServe(":8080",nil)
}

func HomePage(w http.ResponseWriter, r *http.Request)  {
	urlPath := r.URL.Path
	fmt.Fprintf(w, "Hello World %s", urlPath)
}
