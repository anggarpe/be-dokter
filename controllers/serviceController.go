package controllers

import (
	"docApp/models"
	services "docApp/services"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

type ServiceController struct {}
var serviceService services.ServicesService

func (*ServiceController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, serviceService.FindById(id))
}

func (*ServiceController) FindAll(c echo.Context) error {
	res := serviceService.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*ServiceController) Create(c echo.Context) error {
	var services *models.Services
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &services)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Service: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := serviceService.Create(services)
	log.Printf("User created: %#v", services)
	return c.JSON(http.StatusOK, res)
}

func (*ServiceController) Update(c echo.Context) error {
	var services *models.Services
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &services)
	if err != nil{
		log.Printf("Failed Unmarshall in Update Service: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := serviceService.Update(services)
	log.Printf("Service updated: %#v", services)
	return c.JSON(http.StatusOK, res)

}

func (*ServiceController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, serviceService.Delete(id))
}
