package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleMainRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func handleRegister(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	statusCode := userStore.Register(&user)
	c.Status(statusCode)
}

func handleLogin(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	statusCode, userID := userStore.Login(&user)
	c.String(statusCode, userID)
}

func handlePrivateMainRoute(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		log.Println("claims not found")
		return
	}

	jwtClaims, ok := claims.(jwtClaims)
	if !ok {
		log.Println("jwt claims not found")
		return
	}

	c.String(http.StatusOK, jwtClaims.Username)

}
