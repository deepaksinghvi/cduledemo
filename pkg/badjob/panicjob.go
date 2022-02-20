package badjob

import (
	"github.com/deepaksinghvi/cdule/pkg/cdule"
	log "github.com/sirupsen/logrus"
)

type TestPanicJob struct {
	Job cdule.Job
}

func (j TestPanicJob) Execute(jobData map[string]string) {
	log.Info("In MyJPanicJobob")
	a := 0
	i := 100 / a
	log.Infof("i: %d", i)
}

func (j TestPanicJob) JobName() string {
	return "badjob.PanicJob"
}

func (j TestPanicJob) GetJobData() map[string]string {
	return nil
}
