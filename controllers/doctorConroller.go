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

type DoctorController struct {}
var doctorService service.DoctorService

func (*DoctorController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, doctorService.FindById(id))
}

func (*DoctorController) FindAll(c echo.Context) error {
	res := doctorService.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*DoctorController) Create(c echo.Context) error {
	var doctor *models.Doctor
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &doctor)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Doctor: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := doctorService.Create(doctor)
	log.Printf("User created: %#v", doctor)
	return c.JSON(http.StatusOK, res)
}

func (*DoctorController) Update(c echo.Context) error {
	var doctor *models.Doctor
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &doctor)
	if err != nil{
		log.Printf("Failed Unmarshall in Update Doctor: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := doctorService.Update(doctor)
	log.Printf("Doctor updated: %#v", doctor)
	return c.JSON(http.StatusOK, res)

}

func (*DoctorController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, doctorService.Delete(id))
}
