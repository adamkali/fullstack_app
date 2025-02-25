package services

import (
	"github.com/adamkali/fullstack_app/requests"
	"github.com/labstack/echo"
)

type ValidatorService struct{}

func (ValidatorService ValidatorService) ValidateNewUserRequest(e echo.Context) (*requests.NewUserRequest, error) {
    var validRequest requests.NewUserRequest 
    if err := e.Bind(validRequest); err != nil {
        return nil, err
    }
    return &validRequest, nil
}

func (ValidatorService ValidatorService) ValidateLoginRequest(e echo.Context) (*requests.LoginRequest, error) {
    var validRequest requests.LoginRequest
    if err := e.Bind(validRequest); err != nil {
        return nil, err
    }
    return &validRequest, nil
}
