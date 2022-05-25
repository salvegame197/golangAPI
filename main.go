package main

import (
	"salvegame197/golangApi/routes"

	_ "gorm.io/driver/postgres"
)

func main() {

	routes.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	routes.Router.Run("0.0.0.0:3434")

}
