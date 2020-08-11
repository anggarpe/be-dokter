package controllers

import (
	"docApp/models"
	service "docApp/servicess"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

type UserController struct {}
var servUser service.UserService

func (*UserController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servUser.FindById(id))
}

func (*UserController) FindAll(c echo.Context) error {
	res := servUser.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*UserController) Create(c echo.Context) error {
	var user *models.User
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &user)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := servUser.Create(user)
	log.Printf("User created: %#v", user)
	return c.JSON(http.StatusOK, res)
}

func (*UserController) Update(c echo.Context) error {
	var user *models.User
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &user)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := servUser.Update(user)
	log.Printf("User updated: %#v", user)
	return c.JSON(http.StatusOK, res)

}

func (*UserController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servUser.Delete(id))
}



