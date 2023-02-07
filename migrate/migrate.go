package main

import (
	"fmt"
	"github.com/charlie-goldenowl/golangstarter1/initializers"
	"github.com/charlie-goldenowl/golangstarter1/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Book{})
	fmt.Println("? Migration complete")
}
