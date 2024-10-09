package main

import (
	"github.com/gin-gonic/gin"
	"github.com/santhoshkumarr231299/go-gin-crud/config"
	"github.com/santhoshkumarr231299/go-gin-crud/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoutes(router)

	router.Run(":8080")
}
