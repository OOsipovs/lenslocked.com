package main

import (
	"fmt"

	"../models"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "olegsosipovs"
	password = "141082"
	dbname   = "lenslocked"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}

	us.DestructiveReset()
	//defer us.Close()
	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	user.Email = "les.feerdinand@nw.co.uk"
	if err := us.Update(&user); err != nil {
		panic(err)
	}

	userByEmail, err := us.ByEmail("les.feerdinand@nw.co.uk")
	if err != nil {
		panic(err)
	}

	fmt.Println(userByEmail)

	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(userByID)
}
