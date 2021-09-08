package metrics

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func newHighRiskSecurityIssues(config *models.Config) *highRiskSecurityIssues {
	return &highRiskSecurityIssues{
		client: &clients.Asana{
			PersonalAccessToken: config.AsanaPersonalAccessToken,
		},
	}
}

type highRiskSecurityIssues struct {
	client *clients.Asana
}

func (s *highRiskSecurityIssues) Name() string {
	return "high-risk-security-issues"
}

func (s *highRiskSecurityIssues) Value() (float64, error) {
	// TODO - return a count of all tickets that are open and P0 and Type == Security.
	// This will look in Asana for anything in the platform team tagged with "high-risk-security-issue".
	return 0, nil
}
