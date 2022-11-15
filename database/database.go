package database

import (
	"duomly.com/go-bank-backend/helpers"
	"github.com/jinzhu/gorm"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function
func InitDatabase() {
	database, err := gorm.Open("mysql", "sql12570086:AplNvJebrz@tcp(sql12.freesqldatabase.com:3306)/sql12570086?parseTime=True")
	helpers.HandleErr(err)
	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
