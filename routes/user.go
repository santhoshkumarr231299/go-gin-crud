package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santhoshkumarr231299/go-gin-crud/controller"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:name", controller.DeleteUser)
	router.PUT("/:name", controller.UpdateUser)
}
