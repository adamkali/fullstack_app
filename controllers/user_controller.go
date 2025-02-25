package controllers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo"
)

type UserController struct {
    Name string
    DB *pgx.Conn
}

func BuildUserController (p *ControllerParams) UserController{
    return UserController{
        Name: "/user",
        DB: p.DB,
    }
}

func (uc  UserController ) Attatch(e *echo.Echo)  {
    e.Group(uc.Name)
}
