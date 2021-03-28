package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	}
	startServer(port)
}

func startServer(port string) {
	mainRouter := gin.Default()

	mainRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	mainRouter.Run(port)
}
