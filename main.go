package main

import (
	"fibo_api/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.SetRouter(e)
}
