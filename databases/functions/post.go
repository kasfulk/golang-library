package functions

import (
	"github.com/kasfulk/golang-echo-mysql/databases/build"
	"github.com/kasfulk/golang-echo-mysql/databases/schemas"
	"gorm.io/gorm"
)

func ShowPost() []schemas.Post {
	var posts []schemas.Post
	var DB = build.ConnectDatabase()
	DB.Find(&posts)
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
	return posts
}

func ShowPostDetail(id string) (schemas.Post, error) {
	var post schemas.Post
	var DB = build.ConnectDatabase()
	var err error
	if err := DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return post, err
		}
	}
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
	return post, err
}

func DeletePost(id string) int32 {
	var post schemas.Post
	var DB = build.ConnectDatabase()
	RowsAffected := DB.Delete(&post, id).RowsAffected
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
	return int32(RowsAffected)
}
