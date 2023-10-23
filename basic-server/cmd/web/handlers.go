package main

import (
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
