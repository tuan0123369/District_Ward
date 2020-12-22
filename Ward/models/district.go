package models

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Ward struct {
	WardCode string `json:"wardCode"`
	WardName string `json:"wardName"`
}

//Hàm kiểm lỗi
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateWardTable(db *sql.DB) {
	_, err := db.Query("CREATE TABLE ward (wardCode varchar(10) primary key, wardName varchar(50))")
	checkError(err)
}

func InsertTable(db *sql.DB, Code string, Name string) {
	value, e := db.Prepare("INSERT INTO ward(wardCode, wardName) values($1,$2)")
	checkError(e)
	value.Exec(Code, Name)
}

func ShowValues(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT wardcode, wardname FROM ward")
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Story not found",
		})
	}
	post := Ward{}
	for rows.Next() {
		var code, name string

		err = rows.Scan(&code, &name)
		if err != nil {
			panic(err.Error())
		}
		post.WardCode = code
		post.WardName = name
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
