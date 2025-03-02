package responses

import (
	"github.com/adamkali/fullstack_app/internal/repository"
	"github.com/labstack/echo/v4"
)

type LoginResponse struct {
	Data    *repository.User `json:"data"`
	JWT     string           `json:"jwt"`
	Success bool             `json:"success"`
	Message string           `json:"message"`
} // @name LoginResponse

func NewLoginResponse() *LoginResponse {
	return &LoginResponse{Success: false, Message: ""}
}

func (LoginResponse *LoginResponse) Fail(ctx echo.Context, code int, err error) error {
	LoginResponse.Message = err.Error()
	return ctx.JSON(code, LoginResponse)
}

func (LoginResponse *LoginResponse) Successful(ctx echo.Context, user *repository.User, token string) error {
	LoginResponse.Data = user
    LoginResponse.JWT = token
	LoginResponse.Success = true
	return ctx.JSON(200, LoginResponse)
}
