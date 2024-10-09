package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/santhoshkumarr231299/go-gin-crud/config"
	"github.com/santhoshkumarr231299/go-gin-crud/models"
)

func GetUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(200, &users)
}

func CreateUser(ctx *gin.Context) {
	users := models.User{}
	ctx.BindJSON(&users)
	config.DB.Create(&users)
	ctx.JSON(200, &users)
}

func DeleteUser(ctx *gin.Context) {
	users := models.User{}
	config.DB.Where("name = ?", ctx.Param("name")).Delete(&users)
	ctx.JSON(200, &users)
}

func UpdateUser(ctx *gin.Context) {
	users := models.User{}
	config.DB.Where("name = ?", ctx.Param("name")).First(&users)
	ctx.BindJSON(&users)
	config.DB.Save(&users)
	ctx.JSON(200, &users)
}
