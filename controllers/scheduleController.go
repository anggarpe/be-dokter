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

type ScheduleController struct {}
var scheduleService service.ScheduleService

func (*ScheduleController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, scheduleService.FindById(id))
}

func (*ScheduleController) FindAll(c echo.Context) error {
	res := scheduleService.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*ScheduleController) Create(c echo.Context) error {
	var schedule *models.Schedule
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &schedule)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Schedule: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := scheduleService.Create(schedule)
	log.Printf("Schedule created: %#v", schedule)
	return c.JSON(http.StatusOK, res)
}

func (*ScheduleController) Update(c echo.Context) error {
	var schedule *models.Schedule
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &schedule)
	if err != nil{
		log.Printf("Failed Unmarshall in Update Schedule: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := scheduleService.Update(schedule)
	log.Printf("Schedule updated: %#v", schedule)
	return c.JSON(http.StatusOK, res)

}

func (*ScheduleController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, scheduleService.Delete(id))
}
