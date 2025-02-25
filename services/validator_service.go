package services

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"unicode"

	"github.com/adamkali/fullstack_app/requests"
	"github.com/labstack/echo"
)

type ValidatorService struct{}

func validateUsername(username string) bool {
    trimmed := strings.TrimSpace(username)
    if trimmed == "" {
        return false
    }
    pattern := `^[a-zA-Z0-9]+$`
    _, err := regexp.MatchString(pattern, trimmed)
    if err != nil {
        return false
    }
    return true
}

func validatePassword(s string) (sevenOrMore, number, upper, special bool) {
    letters := 0
    for _, c := range s {
        switch {
        case unicode.IsNumber(c):
            number = true
        case unicode.IsUpper(c):
            upper = true
            letters++
        case unicode.IsPunct(c) || unicode.IsSymbol(c):
            special = true
        case unicode.IsLetter(c) || c == ' ':
            letters++
        default:
            //return false, false, false, false
        }
    }
    sevenOrMore = letters > 7
    return
}

func (ValidatorService ValidatorService) ValidateNewUserRequest(e echo.Context) (*requests.NewUserRequest, error) {
    validRequest := new(requests.NewUserRequest)
	if err := e.Bind(&validRequest); err != nil {
		return nil, err
	}

    if !validateUsername(validRequest.Username) {
       return nil, fmt.Errorf("Validation failed (%s) is not a valid username", validRequest.Username) 
    }

    _, err := mail.ParseAddress(validRequest.Email)
    if err != nil {
        return nil, err
    }

    sevenOrMore, number, upper, special := validatePassword(validRequest.Password)
    if  !(sevenOrMore && number && upper && special) {
        return nil, fmt.Errorf(
            "Validation failed. Seven Or More (%t), Number (%t), Upper (%t), Special (%t)",
            sevenOrMore,
            number,
            upper,
            special)
    }

	return validRequest, nil
}

// ValidateLoginRequest validates the Body by using the echo.Context.Bind(requsts.LoginRequest)
// and the returns then last part
func (ValidatorService ValidatorService) ValidateLoginRequest(e echo.Context) (*requests.LoginRequest, error) {
    validRequest := new(requests.LoginRequest)
	if err := e.Bind(&validRequest); err != nil {
		return nil, err
	}
	return validRequest, nil
}
