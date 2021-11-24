package asana

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers/asana/metrics"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func NewMetricProvider(config *models.Config) *MetricProvider {
	return &MetricProvider{
		config:         config,
		bugsThisWeek:   metrics.NewBugsThisWeek(config),
		securityIssues: metrics.NewSecurityIssues(config),
	}
}

type MetricProvider struct {
	config         *models.Config
	bugsThisWeek   *metrics.BugsThisWeek
	securityIssues *metrics.SecurityIssues
}

func (m *MetricProvider) Calculate() ([]*models.Metric, error) {
	var ret []*models.Metric

	calculated, err := m.securityIssues.Calculate()
	if err != nil {
		return nil, err
	}
	ret = append(ret, calculated...)

	calculated, err = m.bugsThisWeek.Calculate()
	if err != nil {
		return nil, err
	}
	ret = append(ret, calculated...)

	return ret, nil
}
