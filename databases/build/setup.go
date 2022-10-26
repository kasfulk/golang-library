package build

import (
	"os"

	"github.com/kasfulk/golang-library/databases/schemas"
	"golang.org/x/crypto/bcrypt"
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

func UserSeed() {
	DB := ConnectDatabase()
	DB.Raw("TRUNCATE TABLE user;")
	password := []byte("Secret")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	var user = schemas.User{
		Username: "admin",
		Password: string(hashedPassword),
		Email:    "admin@gmail.com",
	}
	DB.Create(&user)
}

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
