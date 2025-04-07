package main

import (
	"flag"
	//"log"
	"crypto/tls"
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore" // New import
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"vedanth.snippetbox.net/internal/models"
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

	sessionManager.Cookie.Secure= true

	tlsConfig:= &tls.Config{
		CurvePreferences: []tls.CurveID{
			tls.X25519,
			tls.CurveP256,
		},
	}

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
		users: &models.UserModel{DB: db},
		templateCache: templateCache,
		formDecoder: formDecoder,
		sessionManager: sessionManager,
	}
    srv:= &http.Server{
		Addr: *addr,
		Handler: app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig: tlsConfig,

		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("Starting server on", "port",  srv.Addr)
	err= srv.ListenAndServeTLS("./tls/localhost.pem", "./tls/localhost-key.pem")
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
	users *models.UserModel
	templateCache map[string]*template.Template
	formDecoder *form.Decoder
	sessionManager *scs.SessionManager
}
