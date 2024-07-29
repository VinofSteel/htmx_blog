package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func dbInitializer() *sql.DB {
	var (
		user     string = os.Getenv("PGUSER")
		host     string = os.Getenv("PGHOST")
		password string = os.Getenv("PGPASSWORD")
		port     string = os.Getenv("PGPORT")
		dbName   string = os.Getenv("PGDATABASE")
	)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)

	log.Printf("Opening connection with database %s on port %s...\n", dbName, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB connection:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging db:", err)
	}
	log.Println("Connection opened succesfully!")

	return db
}
