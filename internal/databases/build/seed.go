package build

import (
	"github.com/kasfulk/golang-library/internal/databases/schemas"
	"golang.org/x/crypto/bcrypt"
)

func UserSeed() {
	DB := ConnectDatabase()
	errorExec := DB.Exec("TRUNCATE TABLE users").Error
	if errorExec != nil {
		panic(errorExec)
	}
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
