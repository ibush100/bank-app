package main

import (
	"bank-app/api"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func main() {
	// connectDB()
	// migrate()
	// seed()
	api.StartApi()
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=bankapp sslmode=disable")
	handleErr(err)
	return db
}

func seed() {
	db := connectDB()
	db.Create(&User{Username: "shit bird", Password: "your face"})
	defer db.Close()

}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func migrate() {
	db := connectDB()
	db.AutoMigrate(&User{})
}
