package main

import (
	"fmt"

	"github.com/jonastrindade/go-api/database"
	"github.com/jonastrindade/go-api/models"
	"github.com/jonastrindade/go-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Jonas", Story: "Story 1"},
		{Id: 2, Name: "Jessica", Story: "Story 2"},
	}

	database.ConnectToDatabase()
	fmt.Println("Starting server")
	routes.HandleRequest()
}
