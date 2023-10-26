package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		name := "unknown"
		if len(params.Get("name")) != 0 {
			name = params.Get("name")
		}

		_, _ = fmt.Fprintf(w, "Hi %q", html.EscapeString(name))
	})

	http.ListenAndServe(":8080", nil)
}
