package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err  = gorm.Open("mysql", "root:root@/goproject?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&User{}, &Credential{}, &Ticket{})
	db.Model(&Ticket{}).AddForeignKey("author", "users(email)", "RESTRICT", "RESTRICT")
}
