package job

import (
	"github.com/deepaksinghvi/cdule/pkg/job"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var testJobData map[string]string

type TestJob struct {
	Job job.Job
}

func (j TestJob) Execute(jobData map[string]string) {
	log.Info("In TestJob1")
	for k, v := range jobData {
		valNum, err := strconv.Atoi(v)
		if nil == err {
			jobData[k] = strconv.Itoa(valNum + 1)
		} else {
			log.Error(err)
		}

	}
	testJobData = jobData
	log.Info("In TestJob2")
}

func (j TestJob) JobName() string {
	return "job.TestJob"
}

func (j TestJob) GetJobData() map[string]string {
	return testJobData
}
