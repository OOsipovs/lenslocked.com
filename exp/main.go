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

	//us.DestructiveReset()
	defer us.Close()

	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
