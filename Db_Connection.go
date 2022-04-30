package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

//fmt.Println("Go Mysql Connection..")
//db, err := sql.Open("mysql", "root:aniruddh@tcp(localhost:3306)/mydb")

var urlDSN = "root:aniruddh@tcp(localhost:3306)/mydb?parseTime=true"
var err error

func DataMigration() {
	database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection Failed:(")
	}
	database.AutoMigrate(&Weapons{})

}
