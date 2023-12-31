package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "home.gohtml", data)
}

func (app *application) helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello-world" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World")
}

func (app *application) form(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		app.render(w, r, http.StatusOK, "form.gohtml", app.newTemplateData(r))
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing the form", http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprintf(w, "email: %s \n", r.FormValue("email"))
}
