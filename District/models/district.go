package models

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type District struct {
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
}

//Hàm kiểm lỗi
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//Tạo ra bảng district
func CreateDistrictTable(db *sql.DB, c *gin.Context) {
	_, err := db.Query("CREATE TABLE district (districtCode varchar(10) primary key, districtName varchar(50))")
	if err != nil {
		c.JSON(500, err.Error())
	}
}

//Thêm giá trị vào bảng
func InsertTable(db *sql.DB, Code string, Name string) {
	value, e := db.Prepare("INSERT INTO district(districtCode, districtName) values($1,$2)")
	checkError(e)
	value.Exec(Code, Name)
}

//Hiển thị các record của bảng
func ShowValues(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT districtcode, districtname FROM district")
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Story not found",
		})
	}
	post := District{}

	for rows.Next() {
		var code, name string

		err = rows.Scan(&code, &name)
		if err != nil {
			panic(err.Error())
		}
		post.DistrictCode = code
		post.DistrictName = name
		c.JSON(200, post)
	}
}

//Đọc file csv
func ReadCsv(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	record, err := csv.NewReader(f).ReadAll()
	checkError(err)
	return record
}

//Thêm các phần tử vào table
func InsertValue(db *sql.DB, record [][]string) {
	flag := false
	code := -1
	name := -1
	for _, value := range record {
		if !flag {
			if value[0] == "Name" {
				name = 0
				code = 1
			} else {
				name = 1
				code = 0
			}
			flag = true
		} else {
			InsertTable(db, value[code], value[name])
		}
	}
}
