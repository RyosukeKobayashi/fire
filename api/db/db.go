package db

import (
	"database/sql"
	"fmt"
	"log"

	"api/config"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := config.GetConfig()

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.Postgres.USER, config.Postgres.PASSWORD, config.Postgres.NAME)

	db, _ := sql.Open("postgres", dbinfo)
	defer db.Close()

	// err := db.Ping()
	// if err != nil {
	// 	log.Fatal("Error: Could not establish a connection with the database")
	// }
	DB = db

	CreateUsersTable()
}
