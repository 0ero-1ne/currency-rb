package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"log"
)

type CustomScheduler struct {
	scheduler gocron.Scheduler
}

func NewCustomScheduler() CustomScheduler {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		log.Fatalln("Can not init scheduler: " + err.Error())
	}

	return CustomScheduler{
		scheduler: scheduler,
	}
}

func (scheduler *CustomScheduler) LoadJob(
	jobTime gocron.JobDefinition,
	jobTask gocron.Task,
	options ...gocron.JobOption,
) {
	job, err := scheduler.scheduler.NewJob(jobTime, jobTask, options...)

	if err != nil {
		log.Fatalln("Can not schedule job: " + err.Error())
	}

	log.Printf("Job %s (ID: %v) loaded", job.Name(), job.ID())
}

func (scheduler *CustomScheduler) Start() {
	scheduler.scheduler.Start()
}
