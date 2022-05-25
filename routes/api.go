package routes

import (
	"salvegame197/golangApi/app/http/controllers"
	"salvegame197/golangApi/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var RouterAuth *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(middlewares.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/", controllers.Index)
	}

}
