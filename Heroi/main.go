package main

import (
	"hero-api/dataBase"
	"hero-api/routes"
)

func main() {

	// Connect to the database
	db := dataBase.ConnectDataBase()

	defer db.Close()

	routes.HandleRequests()
}
