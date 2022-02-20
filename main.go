package main

import (
	"github.com/deepaksinghvi/cdule/pkg/cdule"
	"github.com/deepaksinghvi/cdule/pkg/model"
	"github.com/deepaksinghvi/cdule/pkg/utils"
	"github.com/deepaksinghvi/cduledemo/pkg/badjob"
	job2 "github.com/deepaksinghvi/cduledemo/pkg/job"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	c := cdule.Cdule{}
	workerName := os.Args[1:]
	if len(os.Args) > 1 {
		workerName = os.Args[1:]
		log.Infof("Cdule Demo for worker: %s", workerName[0])
		c.NewCduleWithWorker(workerName[0])

	} else {
		c.NewCdule()
	}

	testJob := job2.TestJob{}
	jobData := make(map[string]string)
	jobData["one"] = "1"
	jobData["two"] = "2"
	jobData["three"] = "3"

	testJobModel, err := cdule.NewJob(&testJob, jobData).Build(utils.EveryMinute)
	if nil == testJobModel || nil != err {
		log.Errorf("Job Creation Error for JobName: %s Error: %v", testJob.JobName(), err)
	}

	/*
		This is the job which would create panic and cdule library is able to handle if there are any panics from
		different jobs defined by users.
	*/
	testPanicJob := badjob.TestPanicJob{}
	testPanicJobModel, err := cdule.NewJob(&testPanicJob, nil).Build(utils.EveryMinute)
	if nil == testPanicJobModel || nil != err {
		log.Errorf("Job Creation Error for JobName: %s Error: %v", testPanicJob.JobName(), err)
	}
	time.Sleep(5 * time.Minute)
	c.StopWatcher()

	log.Info("---Cdule Demo---")
}

func printJobs(job *model.Job) {
	log.Infof("JobDetails -> JobName: %s, Schedule Cron: %s, Job Creation Time: %s ", job.JobName, job.CronExpression, job.CreatedAt)
}
