package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nishisuke/ezlog"
)

func main() {
	e := echo.New()

	e.Logger = ezlog.New(os.Stdout, log.DEBUG)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
