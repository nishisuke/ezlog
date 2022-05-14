package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/nishisuke/ezlog"
	"github.com/nishisuke/ezlog/echologger"
)

func main() {
	e := echo.New()

	ezlog.Prepare()
	e.Logger = echologger.New(os.Stdout, log.DEBUG)

	conf := middleware.RequestLoggerConfig{
		LogValuesFunc:    logRequest,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogContentLength: true,
	}
	e.Use(middleware.RequestLoggerWithConfig(conf))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func logRequest(c echo.Context, v middleware.RequestLoggerValues) error {
	c.Logger().Infoj(map[string]interface{}{
		"request": v,
	})
	return nil
}
