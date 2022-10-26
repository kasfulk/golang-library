package main

import (
	"github.com/joho/godotenv"
	"github.com/kasfulk/golang-library/databases/build"
	"github.com/kasfulk/golang-library/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	build.MigrateDatabase()
	routes.BaseRoute()
}
