package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config config
	models models.Models
}

func main() {

	var cfg config
	cfg.port = 9090

	var dsn string

	if dsn = os.Getenv("DSN"); dsn == "" {
		fmt.Println("Not defined")
	} else {
		fmt.Printf("DSN: %s\n", dsn)
	}

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := &application{
		config: cfg,
		models: models.New(db.SQL),
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}
