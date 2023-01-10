package routes

import (
	domain "github.com/kasfulk/golang-library/internal/domain"
)

func BookRoutes() {
	// post route
	e.GET("/book", domain.BookIndex)
	e.POST("/book", domain.BookCreate)
	e.GET("/book/:id", domain.BookDetail)
	e.PUT("/book/:id", domain.BookUpdate)
	e.PATCH("/book/:id", domain.BookedBy)
	e.DELETE("/book/:id", domain.BookDelete)
}
