package metrics

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metrics/util"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

// The last 7 days.
var window = -7 * 24 * time.Hour

func newPlatformStability(config *models.Config) *platformStability {
	return &platformStability{
		config: config,
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
	queryParams := &url.Values{}
	queryParams.Add("teams.any", s.config.AsanaPlatformTeamGid)
	queryParams.Add("completed", "false")
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaTypeFieldGid), "1184099641533292") // Bug.
	queryParams.Add("created_on.after", time.Now().UTC().Add(window).Format(util.YYYYMMDD))

	tasks, err := s.client.SearchTasks(s.config.AsanaWorkspaceGid, *queryParams)
	if err != nil {
		return 0, err
	}
	log.Println("------------- P0 Platform bugs with the last week -------------")
	for _, task := range tasks {
		log.Println(fmt.Sprintf("gid: %s, name: %s", task.GID, task.Name))
	}
	return float64(len(tasks)), nil
}
