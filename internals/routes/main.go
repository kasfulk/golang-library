package routes

import (
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func BaseRoute() {
	IndexRoutes()
	BookRoutes()
	e.Logger.Fatal(e.Start(":1323"))
}
