package config

import (
	"fmt"
	"jin-example/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "123.56.133.187"
	port     = 3306
	user     = "root"
	password = "Lxy15238330738."
	dbName   = "go_test"
	charset  = "utf8mb4"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", user, password, host, port, dbName, charset)
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	helper.ErrorPanic(err)
	return db
}
