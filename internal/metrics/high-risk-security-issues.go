package metrics

import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

func newHighRiskSecurityIssues(config *models.Config) *highRiskSecurityIssues {
	// TODO
	return &highRiskSecurityIssues{}
}

type highRiskSecurityIssues struct {
	// TODO
}

func (s *highRiskSecurityIssues) Name() string {
	return "high-risk-security-issues"
}

func (s *highRiskSecurityIssues) Type() models.MetricType {
	return models.Gauge
}

func (s *highRiskSecurityIssues) Value() (float64, error) {
	// TODO
	return 0, nil
}
