package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/dashboards"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metrics"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func Run(config *models.Config) error {
	allDashboards := dashboards.All(config)
	allMetrics := metrics.All(config)
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case <-ticker.C:
			log.Println("sending metrics")
			if err := runOnce(allDashboards, allMetrics); err != nil {
				log.Println(fmt.Sprintf("error: %s", err))
			}
			log.Println("send finished")
		}
	}
}

func runOnce(allDashboards []models.Dashboard, allMetrics []models.Metric) error {
	for _, dashboard := range allDashboards {
		if err := dashboard.Send(allMetrics); err != nil {
			return err
		}
	}
	return nil
}
