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

type OrderController struct {}
var servOrder service.OrderService

func (*OrderController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servOrder.FindById(id))
}

func (*OrderController) FindAll(c echo.Context) error {
	res := servOrder.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*OrderController) Create(c echo.Context) error {
	var order *models.Order
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &order)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Order: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := servOrder.Create(order)
	log.Printf("Order created: %#v", order)
	return c.JSON(http.StatusOK, res)
}

func (*OrderController) Update(c echo.Context) error {
	var order *models.Order
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &order)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Order: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := servOrder.Update(order)
	log.Printf("Order updated: %#v", order)
	return c.JSON(http.StatusOK, res)

}

func (*OrderController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servOrder.Delete(id))
}



