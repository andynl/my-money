package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

func NewMysqlDatabase() (*gorm.DB, error) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/uangkuu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}