package metrics

import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

func newSecurityCheckCoverage(config *models.Config) *securityCheckCoverage {
	// TODO
	return &securityCheckCoverage{}
}

type securityCheckCoverage struct {
	// TODO
}

func (s *securityCheckCoverage) Name() string {
	return "security-check-coverage"
}

func (s *securityCheckCoverage) Value() (float64, error) {
	// TODO
	// https://docs.github.com/en/rest/reference/apps
	return 0, nil
}
