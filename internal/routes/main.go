package routes

import (
	"github.com/kasfulk/golang-library/internal/config"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func BaseRoute() {
	IndexRoutes()
	BookRoutes()
	e.Logger.Fatal(e.Start(config.ServerPort()))
}
