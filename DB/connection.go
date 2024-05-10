package DB

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectPostgres(dsn string) (*DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		return err
	}
	fmt.Println("Database connected")
	return nil
}
