package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CONFIGURASI MYSQL
func InitMysqlDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("MYSQL_DB_USERNAME"),
		os.Getenv("MYSQL_DB_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
