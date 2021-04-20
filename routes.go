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
	var registerParameters RegisterParameters
	err := c.BindJSON(&registerParameters)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	statusCode := myContext.AuthService.Register(&registerParameters)
	c.Status(statusCode)
}

func handleLogin(c *gin.Context) {
	var loginParameters LoginParameters
	err := c.BindJSON(&loginParameters)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	statusCode, userID := myContext.AuthService.Login(&loginParameters)
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
