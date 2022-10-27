package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kasfulk/golang-library/databases/build"
	dbFunctions "github.com/kasfulk/golang-library/databases/functions"
	"github.com/kasfulk/golang-library/databases/schemas"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookIndex(c echo.Context) error {
	var Redis = build.ConnectRedis()
	redisResult, redisError := Redis.Get("books").Result()

	if redisError != nil && redisResult != "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Kendala pada server",
			"error":   fmt.Sprintf("%v", redisError),
		})
	}

	if redisResult == "" || redisResult == "[]" {
		books, booksError := dbFunctions.ReCacheBooks()
		if booksError != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Kendala pada server",
				"error":   fmt.Sprintf("%v", booksError),
			})
		}
		return c.JSON(http.StatusOK, books)
	} else {
		var resultData []schemas.Book
		err := json.Unmarshal([]byte(redisResult), &resultData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Kendala pada server",
				"error":   fmt.Sprintf("%v", err),
			})
		}
		return c.JSON(http.StatusOK, resultData)
	}

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

func BookCreate(c echo.Context) error {
	book := new(schemas.Book)
	err := c.Bind(book)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Terdapat kesalahan dalam pengiriman data!",
		})
	}

	if err := dbFunctions.CreateBook(book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Terdapat kesalahan dalam pemrosesan data!",
		})
	}
	return c.JSON(http.StatusOK, book)
}

func BookUpdate(c echo.Context) error {
	id := c.Param("id")
	book := new(schemas.Book)
	err := c.Bind(book)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Terdapat kesalahan dalam pengiriman data!",
		})
	}

	if err := dbFunctions.UpdateBook(id, book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Terdapat kesalahan dalam pemrosesan data!",
		})
	}
	return c.JSON(http.StatusOK, book)
}

func BookedBy(c echo.Context) error {
	id := c.Param("id")
	book := new(schemas.Book)
	err := c.Bind(book)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Terdapat kesalahan dalam pengiriman data!",
		})
	}

	if err := dbFunctions.UpdateBook(id, book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Terdapat kesalahan dalam pemrosesan data!",
		})
	}
	return c.JSON(http.StatusOK, book)
}
