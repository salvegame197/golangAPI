package controllers

import (
	"salvegame197/golangApi/app/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	user, count, _ := models.FindUsers(models.User{})

	c.JSON(200, gin.H{
		"message": "Hello World",
		"users":   user,
		"count":   count,
	})
}
