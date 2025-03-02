package responses

import (
	"github.com/adamkali/fullstack_app/internal/repository"
	"github.com/labstack/echo/v4"
)

type UserResponse struct {
    Data *repository.User `json:"data"` 
    Success bool `json:"success"`
    Message string `json:"message"`
} // @name UserResponse

func NewUserResponse() *UserResponse {
    return &UserResponse{ Success: false, Message: "" }
}

func (UserResponse *UserResponse) Fail(ctx echo.Context, code int, err error) error {
    UserResponse.Message = err.Error()
    return ctx.JSON(code, UserResponse)
}

func (UserResponse *UserResponse) Successful(ctx echo.Context, user *repository.User) error {
    UserResponse.Data = user
    UserResponse.Success = true
    return ctx.JSON(200, UserResponse)
}
