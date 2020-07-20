package controllers

import (
	"docApp/models"
	service "docApp/services"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

func FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, service.FindById(id))
}

func GetAllUser(c echo.Context) error {
	res := service.GetAllUser()
	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
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

	res := service.CreateUser(user)
	log.Printf("User created: %#v", user)
	return c.JSON(http.StatusOK, res)
}

func Update(c echo.Context) error {
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

	res := service.UpdateUser(user)
	log.Printf("User updated: %#v", user)
	return c.JSON(http.StatusOK, res)

}

func Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, service.DeleteUser(id))
}



