package internal

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/dashboards"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metrics"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func Run(config *models.Config) error {
	allDashboards := dashboards.All(config)
	allMetrics := metrics.All(config)

	for _, dashboard := range allDashboards {
		if err := dashboard.Send(allMetrics); err != nil {
			return err
		}
	}
	return nil
}
