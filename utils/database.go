package utils

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitializeDatabase() error {
	godotenv.Load()
	db_url := os.Getenv("DB_URL")
	if db_url == ""{
		return errors.New("Incorrect Database URL")
	}
	var err error
	Db, err = sql.Open("postgres", db_url)
	if err != nil{
		panic(err)
	}

	err = Db.Ping()
	if err != nil{
		panic(err)
	}
	log.Print("Database Connected!!")
	return nil
}