package controllers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"

)

type LoginController struct {}
type JwtClaim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func (*LoginController) Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "joko" && password =="jane"{
		//basic cookie
		cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = "some_string_value"
		cookie.Expires = time.Now().Add(12 *time.Hour)
		cookie.Path = "login"

		c.SetCookie(cookie)

		//TODO: create jwt token
		token, err := createJwtToken()
		if err != nil{
			log.Println("Error Create Jwt Token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		//JWT cookie
		JwtCookie := new(http.Cookie)

		JwtCookie.Name = "JwtCookie"
		JwtCookie.Value = token
		JwtCookie.Expires = time.Now().Add(12 *time.Hour)
		JwtCookie.Path = "login"

		c.SetCookie(JwtCookie)

		return c.JSON(http.StatusOK,map[string]string{
			"message":"You were logged in",
			"token":token,
		})
	}

	return c.String(http.StatusUnauthorized, "username or password wrong")
}

func (*LoginController) CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		cookie, err := context.Cookie("sessionID")

		if err != nil{
			log.Print(err)
			if strings.Contains(err.Error(), "named cookie not present"){
				return context.String(http.StatusUnauthorized, "you dont have any cookie")
			}
			return err
		}
		if cookie.Value == "some_string_value"{
			return next(context)
		}
		return context.String(http.StatusUnauthorized, "you dont have the right cookie")
	}
}

func createJwtToken() (string, error){
	claims := JwtClaim{
		"joko",
		jwt.StandardClaims{
			Id: "user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	//this only example, please place key on a file on server or safe place
	//key must be long
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (*LoginController) MainJwt(c echo.Context) error{
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	log.Println("Username: ", claims["name"], "UserId: ", claims["jti"])
	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}
