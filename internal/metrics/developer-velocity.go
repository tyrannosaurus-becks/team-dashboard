package metrics

import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

func newDeveloperVelocity(config *models.Config) *developerVelocity {
	// TODO
	return &developerVelocity{}
}

type developerVelocity struct {
	// TODO
}

func (s *developerVelocity) Name() string {
	return "developer-velocity"
}

func (s *developerVelocity) Type() models.MetricType {
	return models.Gauge
}

func (s *developerVelocity) Value() (float64, error) {
	// TODO
	return 0, nil
}
