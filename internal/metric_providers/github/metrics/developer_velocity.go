package metrics

import (
	"context"
	"strings"
	"time"

	"github.com/google/go-github/v38/github"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
	"golang.org/x/oauth2"
)

const (
	name       = "developer-velocity"
	query      = "org:hipcamp updated:$day_before_yesterday..$yesterday is:pr is:merged"
	timeFormat = "2006-01-02"
)

func NewDeveloperVelocity(config *models.Config) *DeveloperVelocity {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &DeveloperVelocity{
		config:   config,
		ghClient: github.NewClient(tc),
	}
}

type DeveloperVelocity struct {
	config   *models.Config
	ghClient *github.Client
}

func (v *DeveloperVelocity) Calculate() ([]*models.Metric, error) {
	// How many PRs were merged 1-2 days ago?
	now := time.Now().UTC()
	dayBeforeYesterday := now.Add(-48 * time.Hour).Format(timeFormat)
	yesterday := now.Add(-24 * time.Hour).Format(timeFormat)

	q := strings.ReplaceAll(query, "$day_before_yesterday", dayBeforeYesterday)
	q = strings.ReplaceAll(q, "$yesterday", yesterday)
	issueSearchResult, _, err := v.ghClient.Search.Issues(context.Background(), q, nil)
	if err != nil {
		return nil, err
	}

	// Who are our platform devs?
	platformDevs, _, err := v.ghClient.Teams.ListTeamMembersByID(context.Background(), v.config.HipcampOrgID, v.config.PlatformTeamID, nil)
	if err != nil {
		return nil, err
	}

	// How many commits were NOT made by platform developers?
	numCommitsByAppDevs := *issueSearchResult.Total
	for _, issue := range issueSearchResult.Issues {
		if isFromPlatformDeveloper(issue, platformDevs) {
			numCommitsByAppDevs--
		}
	}

	// How many engineers are on the team?
	engineers, _, err := v.ghClient.Teams.ListTeamMembersByID(context.Background(), v.config.HipcampOrgID, v.config.EngineeringTeamID, nil)
	if err != nil {
		return nil, err
	}

	var ret []*models.Metric

	// So then, how many of the team are application developers?
	numAppDevs := len(engineers) - len(platformDevs)
	appDevVelocity := float64(numCommitsByAppDevs) / float64(numAppDevs)
	ret = append(ret, &models.Metric{
		Name:  name,
		Value: appDevVelocity,
		Tags:  &[]string{"application developers"},
	})

	// And what about for the platform devs?
	numPlatformDevs := len(platformDevs)
	numCommitsByPlatformDevs := *issueSearchResult.Total - numCommitsByAppDevs
	platformDevVelocity := float64(numCommitsByPlatformDevs) / float64(numPlatformDevs)
	ret = append(ret, &models.Metric{
		Name:  name,
		Value: platformDevVelocity,
		Tags:  &[]string{"platform developers"},
	})
	return ret, nil
}

func isFromPlatformDeveloper(issue *github.Issue, platformDevs []*github.User) bool {
	author := *issue.User.Login
	for _, platformDev := range platformDevs {
		if *platformDev.Login == author {
			return true
		}
	}
	return false
}
