package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
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
	mainRouter.Use(cors.Default()) //Change later
	mainRouter.GET("/", handleMainRoute)

	mainRouter.POST("/register", handleRegister)
	mainRouter.POST("/login", handleLogin)

	privateRoute := mainRouter.Group("/private")
	privateRoute.Use(requiredAuthentication)
	privateRoute.GET("/", handlePrivateMainRoute)

	mainRouter.Run(port)
}
