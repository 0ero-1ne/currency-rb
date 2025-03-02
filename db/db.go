package db

import (
	"currency/env"
	"currency/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := env.GetEnv(
		"DATABASE_URL",
		"root:@tcp(127.0.0.1:3306)/currency?charset=utf8mb4&parseTime=True&loc=Local",
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	err = database.AutoMigrate(&models.Currency{})

	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	return database
}
