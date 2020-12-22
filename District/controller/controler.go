package controller

import (
	database "../connect"
	"../models"
	"github.com/gin-gonic/gin"
)

//Phương thức GET
func Read(c *gin.Context) {
	db := database.Connect()
	models.ShowValues(db, c)
	defer db.Close()
}

//Phương thức POST
func Insert(c *gin.Context) {
	record := models.ReadCsv("static/data/district.csv")
	db := database.Connect()
	models.CreateDistrictTable(db, c)
	models.InsertValue(db, record)
	c.JSON(200, gin.H{
		"message": "Create and insert success",
	})
	defer db.Close()
}
