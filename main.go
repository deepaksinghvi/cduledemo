package main

import (
	"encoding/json"
	"time"

	"github.com/deepaksinghvi/cdule/pkg/cdule"
	"github.com/deepaksinghvi/cdule/pkg/model"
	"github.com/deepaksinghvi/cdule/pkg/utils"
	"github.com/deepaksinghvi/cduledemo/pkg/badjob"
	job2 "github.com/deepaksinghvi/cduledemo/pkg/job"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Cdule Demo")

	c := cdule.Cdule{}
	c.NewCdule()

	testJob := job2.TestJob{}
	jobData := make(map[string]string)
	jobData["one"] = "1"
	jobData["two"] = "2"
	jobData["three"] = "3"

	testJobModel, err := model.CduleRepos.CduleRepository.GetJobByName(testJob.JobName())
	if nil != err {
		model.CduleRepos.CduleRepository.DeleteJob(testJobModel.ID)
	}

	testJobModel, _ = cdule.NewJob(&testJob, jobData).Build(utils.EveryMinute)
	printJobs(testJobModel)

	/*
		This is the job which would create panic and cdule library is able to handle if there are any panics from
		different jobs defined by users.
	*/
	testPanicJob := badjob.TestPanicJob{}
	testPanicJobModel, err := model.CduleRepos.CduleRepository.GetJobByName(testPanicJob.JobName())
	if nil != err {
		model.CduleRepos.CduleRepository.DeleteJob(testPanicJobModel.ID)
	}
	testPanicJobModel, _ = cdule.NewJob(&testPanicJob, nil).Build(utils.EveryMinute)
	printJobs(testPanicJobModel)

	time.Sleep(5 * time.Minute)
	c.StopWatcher()

	log.Info("---Cdule Demo---")
}

func printJobs(job *model.Job) {
	jsonF, _ := json.Marshal(job)
	log.Info(string(jsonF))
}
