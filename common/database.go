package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToSQL() *sql.DB{
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbParams := os.Getenv("DB_PARAMS")

	log.Printf("connecting to pgsql database=%s", host)
	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s",
		host, port, user, password, dbname, dbParams)

	db, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatalf("conecting to db %s failed: %v", host, err)
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("ping to db %s failed: %v", host, err)
		return nil
	}
	
	log.Printf("success connected to db %s (%s:%s) ...", dbname, host, port)
	return db
}
