package main

import (
	"currency/customscheduler"
	"currency/db"
	"currency/env"
	"currency/models"
	"github.com/gin-gonic/gin"
)

func main() {
	env.Init()
	scheduler := customscheduler.Init()
	server := gin.Default()
	database := db.Init()

	scheduler.Start()

	server.GET("/data", func(c *gin.Context) {
		var currencies []models.Currency
		var day = c.Query("day")

		if len(day) == 0 {
			database.Find(&currencies)
			c.JSON(200, currencies)
			return
		}

		c.JSON(200, day)
	})

	err := server.Run(env.GetEnv("SERVER_PORT", ":8080"))

	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
