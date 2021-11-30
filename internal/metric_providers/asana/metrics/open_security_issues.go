package metrics

import (
	"fmt"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/clients"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

const (
	noPriority = "NoPriority"
	security   = "1200742202724069"
)

func NewSecurityIssues(config *models.Config) *SecurityIssues {
	return &SecurityIssues{
		config: config,
		client: &clients.Asana{
			PersonalAccessToken: config.AsanaPersonalAccessToken,
		},
	}
}

type SecurityIssues struct {
	config *models.Config
	client *clients.Asana
}

func (c *SecurityIssues) Calculate() (ret []*models.Metric, err error) {
	newInLastWeek, err := c.newlyCreated(lastSevenDays, "num-security-issues-opened-last-week")
	if err != nil {
		return nil, err
	}
	ret = append(ret, newInLastWeek...)

	newInLastMonth, err := c.newlyCreated(lastThirtyDays, "num-security-issues-opened-last-month")
	if err != nil {
		return nil, err
	}
	ret = append(ret, newInLastMonth...)

	pastDue, err := c.pastDue()
	if err != nil {
		return nil, err
	}
	ret = append(ret, pastDue...)

	return ret, nil
}

func (c *SecurityIssues) newlyCreated(since time.Duration, metricName string) ([]*models.Metric, error) {
	queryParams := &url.Values{}
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", c.config.AsanaTypeFieldGid), security)
	queryParams.Add("created_on.after", time.Now().UTC().Add(since).Format(iso8601))

	allTasks, err := c.client.SearchTasks(c.config.AsanaWorkspaceGid, *queryParams)
	if err != nil {
		return nil, err
	}
	log.Debugf("for %s, in the last %d days, found %d tasks", metricName, since.Hours()/24, len(allTasks))

	tasksByPriority := map[string][]*clients.ObjectMetadata{
		"P0":       nil,
		"P1":       nil,
		"P2":       nil,
		"P3":       nil,
		noPriority: nil,
	}
	for _, taskMetadata := range allTasks {
		task, err := c.client.GetTask(taskMetadata.GID)
		if err != nil {
			return nil, err
		}
		pxValExists := false
		for _, customField := range task.CustomFields {
			if customField["name"] == "PX" {
				priorityIfc := customField["display_value"]
				if priorityIfc == nil {
					break
				}
				pxValExists = true
				priority := priorityIfc.(string)
				tasksByPriority[priority] = append(tasksByPriority[priority], taskMetadata)
				break
			}
		}
		if !pxValExists {
			tasksByPriority[noPriority] = append(tasksByPriority[noPriority], taskMetadata)
		}
	}

	var ret []*models.Metric
	for px, tasks := range tasksByPriority {
		ret = append(ret, &models.Metric{
			Name:  metricName,
			Value: float64(len(tasks)),
			Tags:  &[]string{px},
		})
	}
	return ret, nil
}

func (c *SecurityIssues) pastDue() ([]*models.Metric, error) {
	queryParams := &url.Values{}
	queryParams.Add(fmt.Sprintf("custom_fields.%s.value", c.config.AsanaTypeFieldGid), security)
	queryParams.Add("completed", "false")
	queryParams.Add("due_on.before", time.Now().UTC().Format(iso8601))

	allTasks, err := c.client.SearchTasks(c.config.AsanaWorkspaceGid, *queryParams)
	if err != nil {
		return nil, err
	}
	return []*models.Metric{
		{
			Name:  "num-security-issues-overdue",
			Value: float64(len(allTasks)),
			Tags:  nil,
		},
	}, nil
}
