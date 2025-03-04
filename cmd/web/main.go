package main

import (
	"flag"
	//"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	
	addr := flag.String("addr", ":4000", "http network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}
	logger.Info("Starting server", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

type application struct {
	logger *slog.Logger
}
