package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexRoutes() {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcom to Library API")
	})
}
