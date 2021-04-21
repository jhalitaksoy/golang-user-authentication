package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var secretKey string

const jwtExpireTime = time.Hour * 24

var myContext *MyContext

func main() {
	LoadEnviromentVariables()

	myContext = newMyContext()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	}

	startServer(port)
}

func LoadEnviromentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error : %v\n", err)
	}

	secretKey = os.Getenv("SECRET_KEY")
	if len(secretKey) == 0 {
		panic("Secret key emty")
	}
}

func CorsConfig() cors.Config {
	return cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true, //Change later
	}
}

func startServer(port string) {
	mainRouter := gin.Default()
	mainRouter.Use(cors.New(CorsConfig()))
	mainRouter.GET("/", handleMainRoute)

	mainRouter.POST("/register", handleRegister)
	mainRouter.POST("/login", handleLogin)

	privateRouter := mainRouter.Group("/private", requiredAuthentication)
	{
		privateRouter.POST("/", handlePrivateMainRoute)
	}

	mainRouter.Run(port)
}
