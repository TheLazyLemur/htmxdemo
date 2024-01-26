package db

import (
	"database/sql"
	_ "embed"
	"log"
)

//go:embed schema.sql
var schema string

func Setup() *sql.DB {
	dbc, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbc.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	return dbc
}
