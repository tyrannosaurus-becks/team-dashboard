package github

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers/github/metrics"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func NewMetricProvider(config *models.Config) *MetricProvider {
	return &MetricProvider{
		config:            config,
		developerVelocity: metrics.NewDeveloperVelocity(config),
	}
}

type MetricProvider struct {
	config            *models.Config
	developerVelocity *metrics.DeveloperVelocity
}

func (m *MetricProvider) Calculate() ([]*models.Metric, error) {
	return m.developerVelocity.Calculate()
}
