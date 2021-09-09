package metrics

import (
	"fmt"
	"net/url"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func NewHighRiskSecurityIssues(config *models.Config) *HighRiskSecurityIssues {
	return &HighRiskSecurityIssues{
		config: config,
		client: &clients.Asana{
			PersonalAccessToken: config.AsanaPersonalAccessToken,
		},
	}
}

type HighRiskSecurityIssues struct {
	config *models.Config
	client *clients.Asana
}

func (s *HighRiskSecurityIssues) Calculate() ([]*models.Metric, error) {
	queryParams := &url.Values{}
	queryParams.Add("teams.any", s.config.AsanaPlatformTeamGid)
	queryParams.Add("completed", "false")
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaTypeFieldGid), "1200742202724069")     // Security.
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaPriorityFieldGid), "1178622795592966") // P0.

	tasks, err := s.client.SearchTasks(s.config.AsanaWorkspaceGid, *queryParams)
	if err != nil {
		return nil, err
	}
	return []*models.Metric{
		{
			Name:  "num-high-risk-security-issues",
			Value: float64(len(tasks)),
			Tags:  nil,
		},
	}, nil
}
