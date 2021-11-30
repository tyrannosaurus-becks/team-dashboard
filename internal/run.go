package internal

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/dashboards"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func Run(config *models.Config) error {
	allDashboards := dashboards.All(config)
	allMetricProviders := metric_providers.All(config)

	// Run at startup.
	log.Info("sending metrics")
	if err := runOnce(allDashboards, allMetricProviders); err != nil {
		log.Error(err)
	}
	log.Info("send finished")

	// Then run every hour.
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case <-ticker.C:
			log.Info("sending metrics")
			if err := runOnce(allDashboards, allMetricProviders); err != nil {
				log.Error(err)
			}
			log.Info("send finished")
		}
	}
}

func runOnce(allDashboards []models.Dashboard, allMetricProviders []models.MetricProvider) error {
	// Retrieve and cache all the metrics to be sent.
	var allMetrics []*models.Metric
	for _, metricProvider := range allMetricProviders {
		metrics, err := metricProvider.Calculate()
		if err != nil {
			return err
		}
		allMetrics = append(allMetrics, metrics...)
	}

	// Send them.
	for _, dashboard := range allDashboards {
		if err := dashboard.Send(allMetrics); err != nil {
			return err
		}
	}
	return nil
}
