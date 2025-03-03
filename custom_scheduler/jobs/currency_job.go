package jobs

import (
	"currency/models"
	"currency/services"
	"encoding/json"
	"github.com/go-co-op/gocron/v2"
	"io"
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

			body, err := io.ReadAll(response.Body)

			if err != nil {
				log.Printf("Can not read response body: " + err.Error())
				return
			}

			var data []*models.Currency

			if err = json.Unmarshal(body, &data); err != nil {
				log.Printf("Can not parse json: " + err.Error())
				return
			}

			currencyService.SaveMany(data)
		},
	)

	return jobDefinition, task
}
