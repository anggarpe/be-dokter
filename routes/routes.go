package routes

import (
	"docApp/services"

	"github.com/labstack/echo"
)

func InitRoute(c echo.Context) error {
	e := c.New()

	e.POST("/users", services.CreateUser)
}
