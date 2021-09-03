package metrics

import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

func newPlatformStability(config *models.Config) *platformStability {
	// TODO
	return &platformStability{}
}

type platformStability struct {
	// TODO
}

func (s *platformStability) Name() string {
	return "platform-stability"
}

func (s *platformStability) Value() (float64, error) {
	// TODO
	return 0, nil
}
