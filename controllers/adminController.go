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

type AdminController struct {}
var adminServ service.AdminService

func (*AdminController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, adminServ.FindById(id))
}

func (*AdminController) FindAll(c echo.Context) error {
	res := adminServ.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*AdminController) Create(c echo.Context) error {
	var user *models.Admin
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = json.Unmarshal(b, &user)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		//return c.String(http.StatusInternalServerError, "")
		return c.JSON(http.StatusInternalServerError, err)
	}


	res := adminServ.Create(user)
	log.Printf("User created: %#v", user)
	return c.JSON(http.StatusOK, res)
}

func (*AdminController) Update(c echo.Context) error {
	var user *models.Admin
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

	res := adminServ.Update(user)
	log.Printf("User updated: %#v", user)
	return c.JSON(http.StatusOK, res)

}

func (*AdminController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, adminServ.Delete(id))
}
