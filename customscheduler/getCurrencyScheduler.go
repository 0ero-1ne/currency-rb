package customscheduler

import (
	"currency/models"
	"encoding/json"
	"github.com/go-co-op/gocron/v2"
	"log"
	"net/http"
	"time"
)

func Init() gocron.Scheduler {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		panic("Can not init customscheduler: " + err.Error())
	}

	_, err = scheduler.NewJob(
		gocron.DurationJob(time.Minute),
		gocron.NewTask(
			func() {
				get, err := http.Get("https://api.nbrb.by/exrates/rates?periodicity=0")
				if err != nil {
					return
				}

				defer get.Body.Close()

				bytes := make([]byte, get.ContentLength)
				_, err = get.Body.Read(bytes)

				var data []models.Currency
				err = json.Unmarshal(bytes, &data)

				if err != nil {
					panic("Can not parse json: " + err.Error())
				}

				log.Printf("data: %+v\n", data[0])
			},
		),
	)

	if err != nil {
		panic("Can not init job: " + err.Error())
	}

	return scheduler
}
