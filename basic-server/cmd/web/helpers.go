package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		fmt.Errorf("%s", err)
		return
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{CurrentYear: time.Now().Year()}
}
