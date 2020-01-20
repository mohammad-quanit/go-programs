package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) { //// here * sign means reading through http request
	fmt.Fprintf(w, "whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) { //// here * sign means reading through http request
	fmt.Fprintf(w, "whoa, this is abput page!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8008", nil)
}
