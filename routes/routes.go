package routes

import (
	"docApp/controllers"
	"github.com/labstack/echo"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.GET("/user", controllers.GetAllUser)
	e.GET("/user/:id", controllers.FindById)
	e.POST("/user", controllers.CreateUser)
	e.DELETE("user/:id", controllers.Delete)
	e.PUT("/user", controllers.Update)

	return e
}
