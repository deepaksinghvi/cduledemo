package main

import (
	"encoding/json"
	"time"

	"github.com/deepaksinghvi/cdule/pkg/cdule"
	"github.com/deepaksinghvi/cdule/pkg/model"
	"github.com/deepaksinghvi/cdule/pkg/utils"
	"github.com/deepaksinghvi/cdule/pkg/watcher"
	"github.com/deepaksinghvi/cduledemo/pkg/badjob"
	job2 "github.com/deepaksinghvi/cduledemo/pkg/job"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Cdule Demo")

	cdule := cdule.Cdule{}
	cdule.NewCdule()

	testJob := job2.TestJob{}
	testJobModel, _ := watcher.NewJob(&testJob, nil).Build(utils.EveryMinute)
	printJobs(testJobModel)

	/*
			This is the job which would create panic and cdule library is able to handle if there are any panics from
		    different jobs defined by users.
	*/
	testPanicJob := badjob.TestPanicJob{}
	testPanicJobModel, _ := watcher.NewJob(&testPanicJob, nil).Build(utils.EveryMinute)
	printJobs(testPanicJobModel)

	time.Sleep(5 * time.Minute)
	cdule.StopWatcher()

	log.Info("---Cdule Demo---")
}

func printJobs(job *model.Job) {
	jsonF, _ := json.Marshal(job)
	log.Info(string(jsonF))
}
