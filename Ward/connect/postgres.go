package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	user         = "admin"
	password     = "1234"
	databaseName = "postgres"
)

func Connect() *sql.DB {
	msql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)
	db, err := sql.Open(databaseName, msql)
	if err != nil {
		panic(err.Error())
	}
	return db
}
