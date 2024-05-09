package models

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 5

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		Product: Product{},
	}
}

type Models struct {
	Product Product
}
