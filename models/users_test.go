package models

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func testingUserService() (*UserService, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "olegsosipovs"
		password = "141082"
		dbname   = "lenslocked_test"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := NewUserService(psqlInfo)

	if err != nil {
		return nil, err
	}
	us.db.LogMode(false)
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}

	user := User{
		Name:  "Alan Shearer",
		Email: "newcastle@nc.co.uk",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Expected ID > 0. Received %d", user.ID)
	}

	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected created at to be recent. Received %s", user.CreatedAt)
	}

	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected updated at to be recent. Received %s", user.UpdatedAt)
	}

}
