package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    //used for routing users CRUD
    router.POST("/user", controllers.CreateUser())//Create a user
    router.GET("/user/:userId", controllers.GetAUser())//Get a user
    router.DELETE("/user/:userId", controllers.DeleteAUser())//Delete a user
    router.GET("/users", controllers.GetAllUsers())//Get all users
}