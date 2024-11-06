package main

import (
	"task-golang-db/routes"
)

func main() {
	router := routes.SetupRoutes()
	router.Run(":8080") // Start the server on port 8080
}
