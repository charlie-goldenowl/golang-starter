package main

import (
	"github.com/charlie-goldenowl/golangstarter1/controllers"
	"github.com/charlie-goldenowl/golangstarter1/initializers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	router := server.Group("/api")

	router.GET("/healthchecker", func(context *gin.Context) {
		message := "Welcome!"
		context.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)

	log.Fatal(server.Run(":" + config.ServerPort))
}
