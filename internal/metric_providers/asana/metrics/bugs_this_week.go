package metrics

import (
	"fmt"
	"net/url"
	"time"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

const name = "num-bugs-opened-this-week"

func NewBugsThisWeek(config *models.Config) *BugsThisWeek {
	return &BugsThisWeek{
		config: config,
		client: &clients.Asana{
			PersonalAccessToken: config.AsanaPersonalAccessToken,
		},
	}
}

type BugsThisWeek struct {
	config *models.Config
	client *clients.Asana
}

func (s *BugsThisWeek) Calculate() ([]*models.Metric, error) {
	// Find all the bugs opened in the last time period.
	queryParams := &url.Values{}
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", s.config.AsanaTypeFieldGid), "1184099641533292") // Bug.
	queryParams.Add("created_on.after", time.Now().UTC().Add(lastSevenDays).Format(iso8601))
	bugs, err := s.client.SearchTasks(s.config.AsanaWorkspaceGid, *queryParams)
	if err != nil {
		return nil, err
	}

	// Get them so we can break them out by team.

	// This map is Project ID : Team Name and is used to cache
	// the teams to whom a particular project belongs, to reduce
	// our load against the Asana API.
	namesForProjectIDs := make(map[string]string)

	// bugsPerTeam will become the metric we emit.
	bugsPerTeam := make(map[string]int)

	for _, bug := range bugs {
		task, err := s.client.GetTask(bug.GID)
		if err != nil {
			return nil, err
		}
		for _, projectMetadata := range task.Projects {
			teamName := ""
			if n, ok := namesForProjectIDs[projectMetadata.GID]; ok {
				teamName = n
			} else {
				project, err := s.client.GetProject(projectMetadata.GID)
				if err != nil {
					return nil, err
				}
				namesForProjectIDs[projectMetadata.GID] = project.Team.Name
				teamName = project.Team.Name
			}
			// Yes, some bugs will be double-counted because of how a task
			// can live in multiple teams, but we don't need it to be perfect,
			// we just need it to give us valuable information we can use to
			// understand the situation.
			bugsPerTeam[teamName]++
		}
	}

	// Give one per team.
	var ret []*models.Metric
	for teamName, numBugs := range bugsPerTeam {
		ret = append(ret, &models.Metric{
			Name:  name,
			Value: float64(numBugs),
			Tags:  &[]string{teamName},
		})
	}
	return ret, nil
}
