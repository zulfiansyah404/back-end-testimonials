package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:Bi1K17ng44YQuOSF0Lxo@tcp(containers-us-west-55.railway.app:5887)/railway?parseTime=true"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Testimonials{})

	DB = database

}
