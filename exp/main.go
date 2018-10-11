package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "yasin"
	dbname = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
	Color string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&User{})

	var u User
	db = db.Where("email = ?", "denssssa@yasin.io").First(&u)
	errors := db.GetErrors()
	if len(errors) > 0 {
		fmt.Println(errors)
		os.Exit(1)
	}

	fmt.Println(u)
}
