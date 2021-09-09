package google_drive

import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

type MetricProvider struct{}

func (m *MetricProvider) Calculate() ([]*models.Metric, error) {
	// TODO
	return nil, nil
}
