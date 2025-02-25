package services_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/adamkali/fullstack_app/requests"
	"github.com/adamkali/fullstack_app/services"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func LoadUserService() (*services.UserService, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err.Error())
	}
	test_context := context.Background()
	POSTGRES_URL := os.Getenv("POSTGRES_URL")
	db, err := pgx.Connect(test_context, POSTGRES_URL)
	if err != nil {
		return nil, err
	}
	return services.CreateUserService(test_context, db), nil
}


func TestCreateUser(t *testing.T) {
	us, err := LoadUserService()
	if err != nil {
		t.Fatalf(`LoadUserService() did not create a user: %v`, err)
	}
	user, err := us.SignUp(&requests.NewUserRequest{
		Username: "testing",
		Email:    "testing@mail.com",
		Password: "testing",
		IsAdmin:  false,
	})
	if err != nil {
		t.Fatalf(`UserService.SignUp(params) did not create a user: %v`, err)
	}
	gotUser, err := us.GetUserById(user.ID)
	if err != nil {
		t.Fatalf(`UserService.GetUserById(%s) did not create a user: %v`, user.ID.String(), err)
	}
	if user.BCryptHash != gotUser.BCryptHash {
		t.Fatalf(`UserService.GetUserById(%s) != UserService.SignUp(params)`, user.ID.String())
	}
    if err = us.RemoveUser(user.ID); err != nil {
        t.Fatalf(`UserService.RemoveUser(%s) did not delete: %v`, user.ID.String(),err)
    }
}

func TestCreateAdminUser(t *testing.T) {
	us, err := LoadUserService()
	if err != nil {
		t.Fatalf(`LoadUserService() did not create a user: %v`, err)
	}
	user, err := us.SignUp(&requests.NewUserRequest{
		Username: "testing2",
		Email:    "testing2@mail.com",
		Password: "testing2",
		IsAdmin:  true,
    })
	if err != nil {
		t.Fatalf(`UserService.SignUp(params) did not create a user: %v`, err)
	}
	gotUser, err := us.GetUserById(user.ID)
	if err != nil {
		t.Fatalf(`UserService.GetUserById(%s) did not create a user: %v`, user.ID.String(), err)
	}
	if user.BCryptHash != gotUser.BCryptHash {
		t.Fatalf(`UserService.GetUserById(%s) != UserService.SignUp(params)`,user.ID.String() )
	}
	if !gotUser.Admin {
		t.Fatalf(`UserService.SignUp(params) did not create an admin user`)
	}
    if err = us.RemoveUser(user.ID); err != nil {
        t.Fatalf(`UserService.RemoveUser(%s) did not delete: %v`, user.ID.String(), err)
    }
}

func TestLoginWithUsername(t *testing.T) {
	us, err := LoadUserService()
	if err != nil {
		t.Fatalf(`LoadUserService() did not initiate: %v`, err)
	}
	user, err := us.SignUp(&requests.NewUserRequest{
		Username: "testing3",
		Email:    "testing3@mail.com",
		Password: "testing3",
		IsAdmin:  true,
	})
	if err != nil {
		t.Fatalf(`UserService.SignUp(params) did not create a user: %v`, err)
	}
	gotUser, err := us.Login(&requests.LoginRequest{
		Username: "testing3",
		Password: "testing3",
	})
	if err != nil {
		t.Fatalf(`UserService.GetUserById(%s) did not create a user: %v`, user.ID.String(), err)
	}
	if user.ID!= gotUser.ID{
		t.Fatalf(`UserService.Login(%s).ID != UserService.SignUp(params).ID`,user.ID.String() )
	}
    if err = us.RemoveUser(user.ID); err != nil {
        t.Fatalf(`UserService.RemoveUser(%s) did not delete: %v`, user.ID.String(), err)
    }
}

func TestLoginWithEmail(t *testing.T) {
	us, err := LoadUserService()
	if err != nil {
		t.Fatalf(`LoadUserService() did not initiate: %v`, err)
	}
	user, err := us.SignUp(&requests.NewUserRequest{
		Username: "testing4",
		Email:    "testing4@mail.com",
		Password: "testing4",
		IsAdmin:  true,
    })
	if err != nil {
		t.Fatalf(`UserService.SignUp(params) did not create a user: %v`, err)
	}
	gotUser, err := us.Login(&requests.LoginRequest{
		Email:    "testing4@mail.com",
		Password: "testing4",
    })
	if err != nil {
		t.Fatalf(`UserService.GetUserById(%s) did not create a user: %v`, user.ID.String(), err)
	}
	if user.ID!= gotUser.ID{
		t.Fatalf(`UserService.Login(%s).ID != UserService.SignUp(params).ID`,user.ID.String() )
	}
    if err = us.RemoveUser(user.ID); err != nil {
        t.Fatalf(`UserService.RemoveUser(%s) did not delete: %v`, user.ID.String(), err)
    }

}
