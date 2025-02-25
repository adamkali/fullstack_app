package controllers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo"
)

type AuthController struct {
    Name string
    DB *pgx.Conn
}

func BuildAuthController (p *ControllerParams) AuthController{
    return AuthController{
        Name: "/auth",
        DB: p.DB,
    }
}

func (ac  AuthController ) Attatch(e *echo.Echo)  {
    e.Group(ac.Name)
}
