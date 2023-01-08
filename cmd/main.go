package main

import (
	"github.com/kasfulk/golang-library/internal/routes"
)

func main() {
	// uncomment this if want to make auto migrate
	// build.MigrateDatabase()
	routes.BaseRoute()
}
