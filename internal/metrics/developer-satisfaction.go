package metrics

import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

func newDeveloperSatisfaction(config *models.Config) *developerSatisfaction {
	// TODO
	return &developerSatisfaction{}
}

type developerSatisfaction struct {
	// TODO
}

func (s *developerSatisfaction) Name() string {
	return "developer-satisfaction"
}

func (s *developerSatisfaction) Type() models.MetricType {
	return models.Gauge
}

func (s *developerSatisfaction) Value() (float64, error) {
	// TODO
	return 0, nil
}
