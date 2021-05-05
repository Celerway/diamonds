package scheduler

import (
	"fmt"
	"github.com/celerway/diamonds/dtos"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func makePrList(prMap dtos.PrMap) string {
	var prList []string
	for pr, link := range prMap {
		prList = append(prList,
			fmt.Sprintf("<%s|[%d]>", link, pr))
	}
	return strings.Join(prList, " ")
}

func (d DiamondSched) dailyReport() {
	daily, err := d.Service.GetStats("day", time.Now().UTC())
	if err != nil {
		log.Errorf("[sched] GetStats(): %v", err)
	}
	var payload []string
	for user, review := range daily {
		payload = append(payload,
			fmt.Sprintf("%s: %s %s", user, review.Badges, makePrList(review.Prs)),
		)
	}

	d.Slapp.Say(strings.Join(payload, "\n"))
	log.Info("[sched] Daily report delivered.")
	time.Sleep(time.Second) // Sleep a second for good measure.
	d.logNextRun()
}

func (d DiamondSched) weeklyReport() {
	weekly, err := d.Service.GetStats("week", time.Now().UTC())
	if err != nil {
		log.Errorf("[sched] GetStats(): %v", err)
	}
	var payload []string
	payload = append(payload, "Rewards awarded for the current week")
	for user, review := range weekly {
		payload = append(payload,
			fmt.Sprintf("%s: %s %s", user, review.Badges, makePrList(review.Prs)),
		)
	}
	d.Slapp.Say(strings.Join(payload, "\n"))
	log.Info("[sched] Weekly report delivered.")
	time.Sleep(time.Second) // Sleep a second for good measure.
	d.logNextRun()
}
