package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDatabase() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3307)/uangku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}
