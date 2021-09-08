package metrics

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

// The last 7 days.
var window = -7 * 24 * time.Hour

func newPlatformStability(config *models.Config) *platformStability {
	return &platformStability{
		client: &clients.Asana{
			PersonalAccessToken: config.AsanaPersonalAccessToken,
		},
	}
}

type platformStability struct {
	config *models.Config
	client *clients.Asana
}

func (s *platformStability) Name() string {
	return "platform-stability"
}

func (s *platformStability) Value() (float64, error) {
	queryParams := url.Values{}
	queryParams.Add("teams.any", s.config.AsanaPlatformTeamGid)
	queryParams.Add("completed", "false")
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaTypeFieldGid), "Bug")
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaPriorityFieldGid), "P0")
	queryParams.Add("created_on.after", time.Now().UTC().Add(window).Format(time.RFC3339))

	tasks, err := s.client.SearchTasks(s.config.AsanaWorkspaceGid, queryParams)
	if err != nil {
		return 0, err
	}
	log.Println("------------- P0 Platform bugs with the last week -------------")
	for _, task := range tasks {
		log.Println(fmt.Sprintf("gid: %s, name: %s", task.GID, task.Name))
	}
	return float64(len(tasks)), nil
}
