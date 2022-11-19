package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {

	options := fmt.Sprintf(
		"%v:@(%v)/%v?charset=%v&parseTime=True&loc=%v",
		"root",
		"localhost",
		"bookstore",
		"utf8mb4",
		"Local",
	)
	d, err := gorm.Open("mysql", "mysql", options)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
