package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"fibo_api/handler"

	"github.com/labstack/echo/v4"
)

// ルーターの設定を書くファイル
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

	e.GET("/fib", handler.FiboHandler)

	err := e.Start(":80")
	return err
}
