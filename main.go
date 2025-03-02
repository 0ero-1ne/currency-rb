package main

import (
	scheduler "currency/custom_scheduler"
	"currency/custom_scheduler/jobs"
	"currency/db"
	"currency/env"
	"currency/services"
	"github.com/gin-gonic/gin"
)

var currencyService = services.NewCurrencyService()

func main() {
	env.Load()
	_ = db.DB()

	customScheduler := scheduler.NewCustomScheduler()
	customScheduler.LoadJob(jobs.NewCurrencyJob())
	customScheduler.Start()

	server := gin.Default()

	server.GET("/data", func(c *gin.Context) {
		day := c.Query("day")

		if len(day) == 0 {
			c.JSON(200, currencyService.FindAll())
			return
		}

		c.JSON(200, currencyService.FindAllByDay(day))
	})

	err := server.Run(env.GetEnv("SERVER_ADDRESS", ":8080"))

	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
