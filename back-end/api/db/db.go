package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // pure Go Postgres driver for the database/sql package
)

//DB do aplicativo ...
type DB struct {
	DB *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mva7155"
	dbname   = "projects"
)

//Connection ...
func (a *DB) Connection() error {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	psqlInfo := os.Getenv("DATABASE_URL")
	var err error
	log.Printf("%s", psqlInfo)

	a.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("[db/OpenConnection] - Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return err
	}
	return nil
}
