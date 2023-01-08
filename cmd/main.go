package main

import (
	"github.com/joho/godotenv"
	"github.com/kasfulk/golang-library/internals/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// build.MigrateDatabase()
	routes.BaseRoute()
}
