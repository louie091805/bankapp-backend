package main

import (
	"duomly.com/go-bank-backend/api"
	"duomly.com/go-bank-backend/database"
)

func main() {
	// Do migration
	// Init database
	database.InitDatabase()
	// migrations.Migrate()
	api.StartApi()
}
