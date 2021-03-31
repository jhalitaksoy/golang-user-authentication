package main

import (
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

}
