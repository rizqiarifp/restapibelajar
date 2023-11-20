package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:8080)/go-restapi-gin"))
if err != nil {
	panic(err)
}
database.AutoMigrate(&Product{})

DB = database
}

