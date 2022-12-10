package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Type Declaration
type user struct {
	Name     string `json : "name" form : "name"`
	Email    string `json : "email" form : "email"`
	Password string `json : "password" form : "password"`
}

// Main function
func main() {
	println("Hello world")
	e := echo.New()

	e.POST("/PostUser", loginPostMethod)

	e.GET("/index", indexGetMethod)

	e.Logger.Info(e.Start(":5555"))
}

// Unexported Function
func loginPostMethod(c echo.Context) error {
	u := new(user)
	fmt.Println("The Content type of the post method is :", c.Request().Header.Get("Content-type"))

	if error := c.Bind(u); error != nil {
		fmt.Println("ini gaad")
		return error
	}
	fmt.Println("Post to /PostUser ", u)
	return c.JSON(http.StatusOK, u)
}

func indexGetMethod(c echo.Context) error {
	return c.String(202, "Hello world from index")
}
