package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/stillsaro/Leenky/internal/database/DBgo"
)

func ConnectToDB(connStr string) (*sql.DB, *DBgo.Queries, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}
	queries := DBgo.New(db)
	log.Println("Successfully connected to DB!")
	return db, queries, nil
}
