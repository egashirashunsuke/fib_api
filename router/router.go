package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo) error {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.GET("/fib", FiboHandler)

	err := e.Start(":80")
	return err
}
