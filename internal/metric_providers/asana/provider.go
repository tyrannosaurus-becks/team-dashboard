package asana

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers/asana/metrics"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func NewMetricProvider(config *models.Config) *MetricProvider {
	return &MetricProvider{
		config:                 config,
		bugsThisWeek:           metrics.NewBugsThisWeek(config),
		highRiskSecurityIssues: metrics.NewHighRiskSecurityIssues(config),
	}
}

type MetricProvider struct {
	config                 *models.Config
	bugsThisWeek           *metrics.BugsThisWeek
	highRiskSecurityIssues *metrics.HighRiskSecurityIssues
}

func (m *MetricProvider) Calculate() ([]*models.Metric, error) {
	var ret []*models.Metric

	calculated, err := m.highRiskSecurityIssues.Calculate()
	if err != nil {
		return nil, err
	}
	ret = append(ret, calculated...)

	calculated, err = m.highRiskSecurityIssues.Calculate()
	if err != nil {
		return nil, err
	}
	ret = append(ret, calculated...)

	return ret, nil
}
