package main

import (
	"bank-app/api"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	api.StartApi()
}
