package controller

import (
	database "../connect"
	"../models"
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	db := database.Connect()
	models.ShowValues(db, c)
	defer db.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}

func Insert(c *gin.Context) {
	record := models.ReadCsv("static/data/ward.csv")
	db := database.Connect()
	models.CreateWardTable(db)
	models.InsertValue(db, record)
	c.JSON(200, gin.H{
		"message": "Create and insert success",
	})
	defer db.Close()
}
