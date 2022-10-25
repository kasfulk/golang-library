package routes

import (
	"github.com/kasfulk/golang-echo-mysql/handlers"
)

func BookRoutes() {
	// post route
	e.GET("/post", handlers.BookIndex)
	e.GET("/post/:id", handlers.BookDetail)
	e.DELETE("/post/:id", handlers.BookDelete)
}
