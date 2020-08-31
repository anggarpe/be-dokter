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

type FacilityController struct {}
var facilityService service.FacilityService

func (*FacilityController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, facilityService.FindById(id))
}

func (*FacilityController) FindAll(c echo.Context) error {
	res := facilityService.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*FacilityController) Create(c echo.Context) error {
	var facility *models.Facility
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &facility)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := facilityService.Create(facility)
	log.Printf("User created: %#v", facility)
	return c.JSON(http.StatusOK, res)
}

func (*FacilityController) Update(c echo.Context) error {
	var facility *models.Facility
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &facility)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := facilityService.Update(facility)
	log.Printf("User updated: %#v", facility)
	return c.JSON(http.StatusOK, res)

}

func (*FacilityController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, facilityService.Delete(id))
}
