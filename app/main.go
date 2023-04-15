package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	http.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is detail page")
	})
	http.ListenAndServe(":8080", nil)
}
