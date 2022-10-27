package functions

import (
	"encoding/json"

	"github.com/kasfulk/golang-library/databases/build"
	"github.com/kasfulk/golang-library/databases/schemas"
	"gorm.io/gorm"
)

func ShowBook() []schemas.Book {
	var books []schemas.Book
	var DB = build.ConnectDatabase()
	DB.Find(&books)
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
	return books
}

func ShowBookDetail(id string) (schemas.Book, error) {
	var book schemas.Book
	var DB = build.ConnectDatabase()
	var err error
	if err := DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return book, err
		}
	}
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
	return book, err
}

func DeleteBook(id string) int32 {
	var book schemas.Book
	var DB = build.ConnectDatabase()
	RowsAffected := DB.Delete(&book, id).RowsAffected
	ReCacheBooks()
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
	return int32(RowsAffected)
}

func CreateBook(schema *schemas.Book) error {
	var DB = build.ConnectDatabase()
	err := DB.Create(schema).Error
	ReCacheBooks()
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		defer stmtManger.Close()
	}
	if err != nil {
		return err
	}
	return err
}

func UpdateBook(id string, schema *schemas.Book) error {
	var DB = build.ConnectDatabase()
	err := DB.Where("id = ?", id).Updates(schema).Error
	ReCacheBooks()
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		defer stmtManger.Close()
	}
	if err != nil {
		return err
	}
	return err
}

func BookedBy(id string, user string) error {
	var DB = build.ConnectDatabase()
	err := DB.Where("id = ?", id).Update("booked_by", user).Error
	ReCacheBooks()
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		defer stmtManger.Close()
	}
	if err != nil {
		return err
	}
	return err
}

func ReCacheBooks() ([]schemas.Book, error) {
	var Redis = build.ConnectRedis()
	books := ShowBook()

	bookJson, bookByteError := json.Marshal(books)
	if bookByteError != nil {
		return nil, bookByteError
	}
	Redis.Set("books", bookJson, 0)
	return books, nil
}
