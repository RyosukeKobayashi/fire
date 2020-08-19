package main

import (
	_ "database/sql"

	"api/db"
	"api/router"

	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()

	r.Run(":8000")
}
