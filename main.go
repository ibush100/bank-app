package main

import (
	"bank-app/api"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// connectDB()
	// migrate()
	// seed()
	api.StartApi()
}
