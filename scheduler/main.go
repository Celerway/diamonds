package scheduler

import (
	"github.com/celerway/diamonds/service"
	"github.com/celerway/diamonds/slapp"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"time"
)

const DAILY = "16:00:00"
const WEEKLY = "23:50:00"

type DiamondSched struct {
	Service  service.DiamondService
	Slapp    slapp.Slapp
	Schedule *gocron.Scheduler
}

func (d DiamondSched) logNextRun() {
	_, nextrun := d.Schedule.NextRun()
	log.Infof("[sched] Next job is at %v", nextrun)
}

func Initialize(srv service.DiamondService, sl slapp.Slapp) DiamondSched {
	d := DiamondSched{
		Service:  srv,
		Slapp:    sl,
		Schedule: gocron.NewScheduler(time.Local),
	}
	s := d.Schedule

	// I don't bother to handle errors on these.
	_, _ = s.Every(1).Monday().At(DAILY).Do(d.dailyReport)
	_, _ = s.Every(1).Tuesday().At(DAILY).Do(d.dailyReport)
	_, _ = s.Every(1).Wednesday().At(DAILY).Do(d.dailyReport)
	_, _ = s.Every(1).Thursday().At(DAILY).Do(d.dailyReport)
	_, _ = s.Every(1).Friday().At(DAILY).Do(d.dailyReport)
	_, _ = s.Every(1).Sunday().At(WEEKLY).Do(d.weeklyReport)

	// When debugging you might wanna have it spit out reports all the time.
	// Make sure to change the channel, though.
	//	s.Every(15).Second().Do(d.dailyReport)

	s.StartAsync() // Spins off a goroutine.
	_, nextrun := s.NextRun()
	log.Infof("[sched] Scheduler initialized. Next job is at %v", nextrun)

	return d
}
