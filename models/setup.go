package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:nlLDFCsjQliGv24k6z1S@tcp(containers-us-west-59.railway.app:7001)/railway?parseTime=true"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Testimonials{});

	DB = database


}
