package migrations

import (
	"duomly.com/go-bank-backend/database"
	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/interfaces"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Refactor createAccounts to use database package
func createAccounts() {
	users := &[2]interfaces.User{
		{Username: "Janica", Email: "jasedillo1207@gmail.com"},
		{Username: "Angela", Email: "ghelparedes@gmail.com"},
	}
	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		database.DB.Create(&account)
	}
}

// Refactor Migrate
func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, &Transactions)

	createAccounts()
}

// Delete Migrate transactions
