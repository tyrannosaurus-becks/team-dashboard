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
	query      = "org:hipcamp updated:$day_before_yesterday..$yesterday is:pr is:merged"
	timeFormat = "2006-01-02"
)

func newDeveloperVelocity(config *models.Config) *developerVelocity {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &developerVelocity{
		config:   config,
		ghClient: github.NewClient(tc),
	}
}

type developerVelocity struct {
	config   *models.Config
	ghClient *github.Client
}

func (s *developerVelocity) Name() string {
	return "developer-velocity"
}

func (s *developerVelocity) Type() models.MetricType {
	return models.Gauge
}

func (s *developerVelocity) Value() (float64, error) {
	// How many PRs were merged 1-2 days ago?
	now := time.Now().UTC()
	dayBeforeYesterday := now.Add(-48 * time.Hour).Format(timeFormat)
	yesterday := now.Add(-24 * time.Hour).Format(timeFormat)

	q := strings.ReplaceAll(query, "$day_before_yesterday", dayBeforeYesterday)
	q = strings.ReplaceAll(q, "$yesterday", yesterday)
	issueSearchResult, _, err := s.ghClient.Search.Issues(context.Background(), "org:hipcamp updated:2021-09-01..2021-09-02 is:pr is:merged", nil)
	if err != nil {
		return 0, err
	}

	// Who are our platform devs?
	platformDevs, _, err := s.ghClient.Teams.ListTeamMembersByID(context.Background(), s.config.HipcampOrgID, s.config.PlatformTeamID, nil)
	if err != nil {
		return 0, err
	}

	// How many commits were NOT made by platform developers?
	numCommitsByAppDevs := len(issueSearchResult.Issues)
	for _, issue := range issueSearchResult.Issues {
		if isFromPlatformDeveloper(issue, platformDevs) {
			numCommitsByAppDevs--
		}
	}

	// How many engineers are on the team?
	engineers, _, err := s.ghClient.Teams.ListTeamMembersByID(context.Background(), s.config.HipcampOrgID, s.config.EngineeringTeamID, nil)
	if err != nil {
		return 0, err
	}

	// So then, how many of the team are application developers?
	numAppDevs := len(engineers) - len(platformDevs)
	velocity := float64(numCommitsByAppDevs) / float64(numAppDevs)
	return velocity, nil
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
