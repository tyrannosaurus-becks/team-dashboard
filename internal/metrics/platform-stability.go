package metrics

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func newPlatformStability(config *models.Config) *platformStability {
	return &platformStability{
		client: &clients.Asana{
			PersonalAccessToken: config.AsanaPersonalAccessToken,
		},
	}
}

type platformStability struct {
	client *clients.Asana
}

func (s *platformStability) Name() string {
	return "platform-stability"
}

func (s *platformStability) Value() (float64, error) {
	// TODO - return a count of all tickets that are Type == Bug and Priority == P0 and open,
	// that were created in the last week.
	return 0, nil
}
