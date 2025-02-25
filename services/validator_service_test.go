package services_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adamkali/fullstack_app/services"
	"github.com/labstack/echo"
)

const (
	newUserRequestJson             = `{"username": "testing1234", "email": "testing1234@mail.com", "password": "superSecret1234!"}`
	newUserRequestJsonInvalidEmail = `{"username": "testing1234", "email": "testing........1234@mail.com", "password": "superSecret1234!"}`
	newUserRequestJsonInvalidUsername = `{"username": "tes!!!!..!!!!ting1234", "email": "testing1234@mail.com", "password": "superSecret1234!"}`
	newUserRequestJsonInvalidPasswordLTS= `{"username": "testing1234", "email": "testing1234@mail.com", "password": "sS!1"}`
	newUserRequestJsonInvalidPasswordUpper= `{"username": "testing1234", "email": "testing1234@mail.com", "password": "super1234!"}`
	newUserRequestJsonInvalidPasswordNumber= `{"username": "testing1234", "email": "testing1234@mail.com", "password": "superSuper!"}`
	newUserRequestJsonInvalidPasswordSpecial= `{"username": "testing1234", "email": "testing1234@mail.com", "password": "superSecret1234"}`
)

func LoadValidatorService() *services.ValidatorService {
	return &services.ValidatorService{}
}

func TestValidNewUserRequestRequest(t *testing.T) {
	validatorService := LoadValidatorService()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserRequestJson))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	_, err := validatorService.ValidateNewUserRequest(c)
	if err != nil {
		t.Fatalf("ValidateLoginRequest(c) could not validate: %v", err)
	}
}

func TestInvalidNewUserRequest_Email(t *testing.T) {
	validatorService := LoadValidatorService()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserRequestJsonInvalidEmail))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_, err := validatorService.ValidateNewUserRequest(c)
	if err == nil {
		t.Fatalf("ValidateLoginRequest(c) did not catch that email was invalid: %v", newUserRequestJson)
	}
}

func TestInvalidLoginRequest_Username(t *testing.T) {
	validatorService := LoadValidatorService()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserRequestJsonInvalidUsername))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_, err := validatorService.ValidateLoginRequest(c)
	if err == nil {
		t.Fatalf("ValidateLoginRequest(c) did not catch that email was invalid: %v", newUserRequestJson)
	}
}

func TestInvalidLoginRequest_Password_LessThanSeven(t *testing.T) {
	validatorService := LoadValidatorService()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserRequestJsonInvalidPasswordLTS))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_, err := validatorService.ValidateLoginRequest(c)
	if err == nil {
		t.Fatalf("ValidateLoginRequest(c) did not catch that email was invalid: %v", newUserRequestJson)
	}
}

func TestInvalidLoginRequest_Password_Upper(t *testing.T) {
	validatorService := LoadValidatorService()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserRequestJsonInvalidPasswordUpper))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_, err := validatorService.ValidateLoginRequest(c)
	if err == nil {
		t.Fatalf("ValidateLoginRequest(c) did not catch that email was invalid: %v", newUserRequestJson)
	}
}

func TestInvalidLoginRequest_Password_Number(t *testing.T) {
	validatorService := LoadValidatorService()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserRequestJsonInvalidPasswordNumber))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_, err := validatorService.ValidateLoginRequest(c)
	if err == nil {
		t.Fatalf("ValidateLoginRequest(c) did not catch that email was invalid: %v", newUserRequestJson)
	}
}

func TestInvalidLoginRequest_Password_Special(t *testing.T) {
	validatorService := LoadValidatorService()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserRequestJsonInvalidPasswordSpecial))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_, err := validatorService.ValidateLoginRequest(c)
	if err == nil {
		t.Fatalf("ValidateLoginRequest(c) did not catch that email was invalid: %v", newUserRequestJson)
	}
}
