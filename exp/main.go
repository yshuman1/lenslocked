package main

import (
	"fmt"

	"lenslocked.com/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	// user.Email = "michael@michaelscottpaperco.com"
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }
	// userByEmail, err := us.ByEmail("michael@michaelscottpaperco.com")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userByEmail)
	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)
}
