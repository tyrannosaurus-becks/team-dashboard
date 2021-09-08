package metrics

import (
	"fmt"
	"log"
	"net/url"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func newHighRiskSecurityIssues(config *models.Config) *highRiskSecurityIssues {
	return &highRiskSecurityIssues{
		config: config,
		client: &clients.Asana{
			PersonalAccessToken: config.AsanaPersonalAccessToken,
		},
	}
}

type highRiskSecurityIssues struct {
	config *models.Config
	client *clients.Asana
}

func (s *highRiskSecurityIssues) Name() string {
	return "high-risk-security-issues"
}

func (s *highRiskSecurityIssues) Value() (float64, error) {
	queryParams := &url.Values{}
	queryParams.Add("teams.any", s.config.AsanaPlatformTeamGid)
	queryParams.Add("completed", "false")
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaTypeFieldGid), "1200742202724069") // Security.
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaPriorityFieldGid), "1178622795592966") // P0.

	tasks, err := s.client.SearchTasks(s.config.AsanaWorkspaceGid, *queryParams)
	if err != nil {
		return 0, err
	}
	log.Println("------------- P0 Platform security issues -------------")
	for _, task := range tasks {
		log.Println(fmt.Sprintf("gid: %s, name: %s", task.GID, task.Name))
	}
	return float64(len(tasks)), nil
}
