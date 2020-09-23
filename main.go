package main

import (
	"bank-app/api"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/* TODO
redo if/else to return happy path and have conditional handle non - error early
*/
func main() {
	// connectDB()
	// migrate()
	// seed()
	api.StartApi()
}
