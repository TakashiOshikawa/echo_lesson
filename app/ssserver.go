package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"./dao/designers"
)

func main() {
	e := echo.New()

	d := designers.FindById(1)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, d)
	})

	e.Run(standard.New(":1323"))
}