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

type ProductController struct {}
var productService service.ProductService

func (*ProductController) FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, productService.FindById(id))
}

func (*ProductController) FindAll(c echo.Context) error {
	res := productService.FindAll()
	return c.JSON(http.StatusOK, res)
}

func (*ProductController) Create(c echo.Context) error {
	var product *models.Product
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &product)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Product: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := productService.Create(product)
	log.Printf("Product created: %#v", product)
	return c.JSON(http.StatusOK, res)
}

func (*ProductController) Update(c echo.Context) error {
	var product *models.Product
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &product)
	if err != nil{
		log.Printf("Failed Unmarshall in Create Product: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	res := productService.Update(product)
	log.Printf("Product updated: %#v", product)
	return c.JSON(http.StatusOK, res)

}

func (*ProductController) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, productService.Delete(id))
}

func (*ProductController) Search(c echo.Context) error {
	name := c.Param("name")
	return c.JSON(http.StatusOK, productService.Search(name))
}
