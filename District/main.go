package main

import (
	controller "./controller"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.GET("/story/Read", controller.Read)
		client.POST("/story/Insert/0202", controller.Insert)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
