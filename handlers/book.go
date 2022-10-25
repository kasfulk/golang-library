package handlers

import (
	"net/http"

	dbFunctions "github.com/kasfulk/golang-echo-mysql/databases/functions"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookIndex(c echo.Context) error {
	posts := dbFunctions.ShowBook()
	return c.JSON(http.StatusOK, posts)
}

func BookDetail(c echo.Context) error {
	id := c.Param("id")
	post, err := dbFunctions.ShowBookDetail(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Data tidak ditemukan!",
			})
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, post)
}

func BookDelete(c echo.Context) error {
	id := c.Param("id")
	RowsAffected := dbFunctions.DeleteBook(id)
	if RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Data yang dihapus tidak ditemukan!",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Data telah dihapus",
	})
}

// func BookCreate(c echo.Context) error {

// }

// func BookUpdate(c echo.Context) error {

// }
