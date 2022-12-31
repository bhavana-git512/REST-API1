package database

import (
	"log"
	"manger/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connectme() {

	var err error

	db, err = gorm.Open("mysql", "root:Bhavana@1998@tcp(127.0.0.1:3306)/code?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Unable to connect to database")
	} else {
		log.Println("Connection Successful")
	}

	db.AutoMigrate(&model.Idlibatch{})

}
