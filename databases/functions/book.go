package functions

import (
	"github.com/kasfulk/golang-echo-mysql/databases/build"
	"github.com/kasfulk/golang-echo-mysql/databases/schemas"
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
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
	return int32(RowsAffected)
}