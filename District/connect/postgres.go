package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost" //host của sever
	port         = 5432        //cổng của host
	user         = "admin"     //Tên người dùng postgres
	password     = "1234"      //Mật khẩu người dùng
	databaseName = "postgres"  //Tên của dataBase
)

//Kết nối tới database
func Connect() *sql.DB {
	msql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)
	db, err := sql.Open(databaseName, msql)
	if err != nil {
		panic(err.Error())
	}
	return db
}
