package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitializeDatabase() error {
	err_env := godotenv.Load()
	if err_env != nil{
		fmt.Print("Error loading .env file")
	}
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