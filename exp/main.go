package main

import (
	"fmt"

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
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Color  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
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
	db.AutoMigrate(&User{}, &Order{})

	var u User
	if err := db.Preload("Orders").Where("email=?, email@yasin.io").First(&u).Error; err != nil {
		panic(err)
	}
	fmt.Println(u)
	fmt.Println(u.Orders)
	createOrder(db, u, 1001, "fake desc #1")
	createOrder(db, u, 9999, "fake desc #2")
	createOrder(db, u, 125, "fake desc #3")
	createOrder(db, u, 99099, "fake desc #4")
	createOrder(db, u, 100, "fake desc #5")

}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}
}
