package main

import (
	"flag"
	//"log"
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
	"time"
	"os"


	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"vedanth.snippetbox.net/internal/models"
	"github.com/alexedwards/scs/mysqlstore" // New import
    "github.com/alexedwards/scs/v2"
)

func main() {
	
	addr := flag.String("addr", ":4000", "http network address")
    
    dsn:= flag.String("dsn", "web:vedanthN@25@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    
	db, err:= openDB(*dsn)

	if err!=nil{
		logger.Error(err.Error())
		os.Exit(1)
	}
    //always close the database connection
	defer db.Close()

	//initialise template cache
	templateCache, err:= newTemplateCache()
	if err!=nil{
		logger.Error(err.Error())
		os.Exit(1)
	}
    
	formDecoder:= form.NewDecoder()
	sessionManager:= scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder: formDecoder,
		sessionManager: sessionManager,
	}
	logger.Info("Starting server", slog.String("addr", *addr))

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil,err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

type application struct {
	logger *slog.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder *form.Decoder
	sessionManager *scs.SessionManager
}
