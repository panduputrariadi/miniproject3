package config

import (
	"fmt"
	"log"
	"os"
	"panduputra/miniproject3/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

func OpenDB() {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	mysqlconn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	Mysql = MysqlDB{
		DB: mysqlconn,
	}

	err = autoMigrate(mysqlconn)
	if err != nil {
		log.Fatalf("failed to auto migrate database: %v", err)
	}
}


func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Book{},
	)
	if err != nil {
		return err
	}

	return nil
}
