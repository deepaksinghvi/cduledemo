package badjob

import (
	"github.com/deepaksinghvi/cdule/pkg/job"
	log "github.com/sirupsen/logrus"
)

type TestPanicJob struct {
	Job job.Job
}

func (j TestPanicJob) Execute() {
	log.Info("In MyJPanicJobob")
	a := 0
	i := 100 / a
	log.Infof("i: %d", i)
}

func (j TestPanicJob) JobName() string {
	return "badjob.PanicJob"
}
