package controllers

import "github.com/labstack/echo"

type Controller interface {
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
	Create(c echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
