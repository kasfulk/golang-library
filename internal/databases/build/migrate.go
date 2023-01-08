package build

import (
	"github.com/kasfulk/golang-library/internal/databases/schemas"
	"gorm.io/gorm"
)

func MigrateDatabase() {
	var db = ConnectDatabase()
	db.AutoMigrate(&schemas.Book{})
	db.AutoMigrate(&schemas.User{})
	UserSeed()
	var stmtManger, ok = db.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
}
