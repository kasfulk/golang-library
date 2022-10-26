package routes

import (
	"github.com/kasfulk/golang-library/handlers"
)

func BookRoutes() {
	// post route
	e.GET("/book", handlers.BookIndex)
	e.POST("/book", handlers.BookCreate)
	e.GET("/book/:id", handlers.BookDetail)
	e.DELETE("/book/:id", handlers.BookDelete)
}
