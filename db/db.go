package db

import (
	"currency/env"
	"currency/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	database   *gorm.DB
	defaultDSN = "root:@tcp(127.0.0.1:3306)/currency?charset=utf8mb4&parseTime=True&loc=Local"
)

func DB() *gorm.DB {
	if database == nil {
		initDB()
		migrateDB()
	}

	return database
}

func initDB() {
	dsn := env.GetEnv("DATABASE_URL", defaultDSN)

	var err error
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database: " + err.Error())
	}
}

func migrateDB() {
	err := database.AutoMigrate(&models.Currency{})

	if err != nil {
		log.Fatalln("Failed to migrate database: " + err.Error())
	}
}
