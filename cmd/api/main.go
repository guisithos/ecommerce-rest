package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/guisithos/ecommerce-rest/DB"
	"github.com/guisithos/ecommerce-rest/internals/models"
)

type config struct {
	port int
}

type application struct {
	config   config
	models   models.Models
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	var cfg config
	cfg.port = 8080

	var dsn string

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if dsn = os.Getenv("DSN"); dsn == "" {
		fmt.Println("DSN not defined.")
	} else {
		fmt.Printf("DSN: %s\n", dsn)
	}

	db, err := DB.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Connection failed")
	}
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		models:   models.New(db.SQL),
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}

func (app *application) serve() error {
	app.infoLog.Println("port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}
