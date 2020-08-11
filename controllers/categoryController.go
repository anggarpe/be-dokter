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

type CategoryController struct {}
var categoryService service.CategoryService

func (*CategoryController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, categoryService.FindById(id))
}

func (*CategoryController) FindAll(c echo.Context) error {
	res := categoryService.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*CategoryController) Create(c echo.Context) error {
	var category *models.Category
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &category)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := categoryService.Create(category)
	log.Printf("User created: %#v", category)
	return c.JSON(http.StatusOK, res)
}

func (*CategoryController) Update(c echo.Context) error {
	var category *models.Category
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &category)
	if err != nil{
		log.Printf("Failed Unmarshall in Create User: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := categoryService.Update(category)
	log.Printf("User updated: %#v", category)
	return c.JSON(http.StatusOK, res)

}

func (*CategoryController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, categoryService.Delete(id))
}
