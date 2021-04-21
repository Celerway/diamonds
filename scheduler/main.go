package scheduler

import (
	"github.com/celerway/diamonds/service"
	"github.com/celerway/diamonds/slapp"
	log "github.com/sirupsen/logrus"
	"time"
)

type DiamondSched struct {
	Service service.DiamondService
	Slapp   slapp.Slapp
}

func Initialize(s service.DiamondService, sl slapp.Slapp) DiamondSched {
	d := DiamondSched{
		Service: s,
		Slapp:   sl,
	}
	log.Info("Scheduler initialized")
	return d
}

func workerActual() {
	log.Info("Scheduler goroutine running")
	for {
		time.Sleep(1 * time.Second)
	}
}

func (ds DiamondSched) Worker() {
	go workerActual()
}
