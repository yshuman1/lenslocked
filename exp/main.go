package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yshuman1/lenslocked.com/models"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "yasin"
	dbname = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	// user, err := us.ByID(1)
	//fmt.Println(user)

}
