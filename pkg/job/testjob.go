package job

import (
	"github.com/deepaksinghvi/cdule/pkg/job"
	log "github.com/sirupsen/logrus"
)

type TestJob struct {
	Job job.Job
}

func (j TestJob) Execute() {
	log.Info("In TestJob1")
	log.Info("In TestJob2")
}

func (j TestJob) JobName() string {
	return "job.TestJob"
}
