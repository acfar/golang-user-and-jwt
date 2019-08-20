package database

import (
"database/sql"
"log"
_"github.com/go-sql-driver/mysql"
)

type Connection struct{}

func (c Connection) Connect() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/news")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
