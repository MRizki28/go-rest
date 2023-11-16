package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_go_rest"))
	if err != nil {
		panic(err)
	}

	DB = database
}