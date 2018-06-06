package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
	"os"
)

var db *gorm.DB

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println("Environnement variable [" + key + "] must be set")
		os.Exit(1)
	}
	return value
}

func databaseAddress() string {
	var address string
	dbUser := getEnv("DATABASE_USER")
	dbPassword := getEnv("DATABASE_PASSWORD")
	dbHost :=getEnv("DATABASE_HOST")
	dbPort := getEnv("DATABASE_PORT")
	dbName := getEnv("DATABASE_NAME")

	address = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	return address
}

func InitDB() {
	var err error

	db, err  = gorm.Open("mysql", databaseAddress())
	if err != nil {
		log.Panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&User{}, &Credential{}, &Ticket{}, &Message{}, &TicketArchive{}, &MessageArchive{})
	db.Model(&Ticket{}).AddForeignKey("author", "users(email)", "RESTRICT", "RESTRICT")
	db.Model(&Message{}).AddForeignKey("ticket", "tickets(id)", "RESTRICT", "RESTRICT")
	db.Model(&Message{}).AddForeignKey("author", "users(email)", "RESTRICT", "RESTRICT")

}
