package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)

	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello World")
}
