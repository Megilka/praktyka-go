package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/miau", miau)
	e.GET("/path/:id/file", file_manager)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func miau(c echo.Context) error {
	msg := fmt.Sprintf("Miauu ~%v\nTyp -> %v <3", c.QueryParam("name"), c.QueryParam("type"))
	return c.String(http.StatusOK, msg)
}

func file_manager(c echo.Context) error {
	id := c.Param("id")
	max_size, _ := strconv.Atoi(c.QueryParam("max_size"))
	var msg string

	if id == "qwerty" {
		msg = fmt.Sprintf("Znaleziono plik qwerty")
	} else if id == "q" {
		msg = fmt.Sprintf("Znaleziono plik q")
	} else {
		msg = fmt.Sprintf("Nie znaleziono pliku %v", id)
	}

	return c.String(http.StatusOK, msg[:min(max_size, len(msg))])
}
