package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

const PORT = ":8080"

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", PORT, "HTTP network address")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	app := application{logger: logger, templateCache: templateCache}

	logger.Info("Starting server on ", "addr", *addr)
	err = http.ListenAndServe(PORT, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
