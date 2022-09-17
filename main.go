package main

import (
	"gin-mongo-api/configs"
	"github.com/gin-gonic/gin"
	"gin-mongo-api/routes"
)

func main() {
	// Set the router as the default one shipped with Gin
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//run mongoDB
	configs.ConnectDB()

	//run routes
	routes.UserRoute(r)

	r.Run("localhost:6000") // listen and serve on

	// Serve frontend static files
	//r.Static("/", "./frontend/dist")
}