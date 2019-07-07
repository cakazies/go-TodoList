package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	fmt.Println("init db running")
	if e := godotenv.Load(); e != nil {
		fmt.Print(e)
	}
	db_username := os.Getenv("db_username")
	db_password := os.Getenv("db_password")
	db_host := os.Getenv("db_host")
	db_name := os.Getenv("db_name")
	db_port := os.Getenv("db_port")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name)
	fmt.Println(dbUri)
	conn, err := gorm.Open("mysql", dbUri)

	if err != nil {
		fmt.Print(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}

func Migrate() {
	db.DropTableIfExists(&Todo{}, &Account{})
	db.Debug().AutoMigrate(&Todo{}, &Account{})
}
