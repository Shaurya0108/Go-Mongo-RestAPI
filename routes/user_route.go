package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    //used for routing users
    router.POST("/user", controllers.CreateUser())
    router.GET("/user/:userId", controllers.GetAUser())
}