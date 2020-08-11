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

type ReservationController struct {}
var servReservation service.ReservationService

func (*ReservationController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servReservation.FindById(id))
}

func (*ReservationController) FindAll(c echo.Context) error {
	res := servReservation.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*ReservationController) Create(c echo.Context) error {
	var reservation *models.Reservation
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &reservation)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Reservation: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := servReservation.Create(reservation)
	log.Printf("Reservation created: %#v", reservation)
	return c.JSON(http.StatusOK, res)
}

func (*ReservationController) Update(c echo.Context) error {
	var reservation *models.Reservation
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &reservation)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Reservation: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := servReservation.Update(reservation)
	log.Printf("Reservation updated: %#v", reservation)
	return c.JSON(http.StatusOK, res)

}

func (*ReservationController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servReservation.Delete(id))
}



