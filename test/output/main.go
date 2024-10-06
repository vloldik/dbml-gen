package main

import (
	"output/migrates"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:test@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))

	if err := migrates.MigrateAll(db); err != nil {
		panic(err)
	}
}
