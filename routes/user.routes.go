package routes

import (
	"github.com/ddld93/jobinn/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.POST("/user", controllers.CreateUser())
        api.GET("/user/:userId", controllers.GetAUser())
        api.PATCH("/user/:userId", controllers.EditAUser())
        api.DELETE("/user/:userId", controllers.DeleteAUser())
        api.GET("/users", controllers.GetAllUsers())

    }
}