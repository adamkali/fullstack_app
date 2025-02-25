package controllers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo"
)

type ControllerParams struct {
    DB *pgx.Conn
}

type IController interface {
    Attatch(e *echo.Echo) 
}

func AttatchControllers(e *echo.Echo, conts ...IController) {
    for _,v := range conts {
        v.Attatch(e)
    }
}

