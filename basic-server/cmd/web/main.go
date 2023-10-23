package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

const PORT = ":8080"

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", PORT, "HTTP network address")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := application{logger: logger}

	logger.Info("Starting server on ", "addr", *addr)
	err := http.ListenAndServe(PORT, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
