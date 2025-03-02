package jobs

import (
	"currency/models"
	"currency/services"
	"encoding/json"
	"github.com/go-co-op/gocron/v2"
	"log"
	"net/http"
)

var currencyService = services.NewCurrencyService()

func NewCurrencyJob() (gocron.JobDefinition, gocron.Task) {
	jobDefinition := gocron.DailyJob(
		1,
		gocron.NewAtTimes(gocron.NewAtTime(12, 00, 0)),
	)

	task := gocron.NewTask(
		func() {
			response, err := http.Get("https://api.nbrb.by/exrates/rates?periodicity=0")

			if err != nil || response.StatusCode > 299 {
				log.Printf("Can not connect to service")
				return
			}

			defer response.Body.Close()

			bytes := make([]byte, response.ContentLength)
			_, err = response.Body.Read(bytes)

			var data []*models.Currency
			err = json.Unmarshal(bytes, &data)

			if err != nil {
				log.Printf("Can not parse json: " + err.Error())
				return
			}

			currencyService.SaveMany(data)
		},
	)

	return jobDefinition, task
}
