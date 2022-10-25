package build

import (
	"os"

	"github.com/kasfulk/golang-echo-mysql/databases/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var err error

	var DBUser = os.Getenv("DB_USERNAME")
	var DBPass = os.Getenv("DB_PASSWORD")
	var DBHost = os.Getenv("DB_HOST")
	var DBPort = os.Getenv("DB_PORT")
	var DBName = os.Getenv("DB_NAME")

	ConnectionString := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8"
	db, err := gorm.Open(mysql.Open(ConnectionString), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	return db
}

func MigrateDatabase() {
	var db = ConnectDatabase()
	db.AutoMigrate(&schemas.Post{})
	var stmtManger, ok = db.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		stmtManger.Close()
	} else {
		panic(stmtManger)
	}
}
