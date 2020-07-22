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

type VarietyController struct {}
var varietyService service.VarietyService

func (*VarietyController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, varietyService.FindById(id))
}

func (*VarietyController) FindAll(c echo.Context) error {
	res := varietyService.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*VarietyController) Create(c echo.Context) error {
	var variety *models.Variety
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &variety)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := varietyService.Create(variety)
	log.Printf("User created: %#v", variety)
	return c.JSON(http.StatusOK, res)
}

func (*VarietyController) Update(c echo.Context) error {
	var variety *models.Variety
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &variety)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := varietyService.Update(variety)
	log.Printf("User updated: %#v", variety)
	return c.JSON(http.StatusOK, res)

}

func (*VarietyController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, varietyService.Delete(id))
}
