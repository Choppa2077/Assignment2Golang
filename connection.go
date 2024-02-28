package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func InitDB() {
	dsn := "host=localhost user=postgres password=m511 dbname=gormcrud port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db

	// AutoMigrate tables
	DB.AutoMigrate(&Student{}, &Course{}, &Department{}, &Enrollment{}, &Instructor{})
}
