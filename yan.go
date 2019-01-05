package main

import (
	"Go/app/Http/Controllers"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/a":
			fmt.Fprintf(w, "Hello, %q", Controllers.Name())
		case "/b":
			fmt.Fprintf(w, "Hello, %q", html.EscapeString("b"))
		default:
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		}
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		i, nil := fmt.Fprintf(w, "Hello, %q", r.URL.Path)
		log.Println(i, nil)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
