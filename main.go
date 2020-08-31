package main

import (
	"docApp/script"
	"docApp/routes"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	_, err := script.DbConn()
	if err != nil{
		os.Exit(1)
	}
	e := routes.InitRoute()
	fmt.Println("Starting Server...")


	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.PUT, echo.GET, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	e.Logger.Fatal(e.Start(":8000"))
}