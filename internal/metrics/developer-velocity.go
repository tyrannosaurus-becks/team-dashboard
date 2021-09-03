package metrics

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/v38/github"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
	"golang.org/x/oauth2"
)

const (
	query      = "author:$username created:$day_before_yesterday..$yesterday"
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
	// This will look in Github for the number of PRs per day from members
	// of the engineering team who are NOT on the platform engineering team.
	appDevs, err := s.getApplicationDevelopers()
	if err != nil {
		return 0, err
	}

	now := time.Now().UTC()
	dayBeforeYesterday := now.Add(-24 * time.Hour).Format(timeFormat)
	yesterday := now.Add(-24 * time.Hour).Format(timeFormat)

	totalCommits := 0
	for _, appDev := range appDevs {
		numCommitsForDev, err := s.numCommitsToMain(appDev, dayBeforeYesterday, yesterday)
		if err != nil {
			return 0, err
		}
		totalCommits += numCommitsForDev
	}
	velocity := float64(totalCommits) / float64(len(appDevs))
	return velocity, nil
}

func (s *developerVelocity) getApplicationDevelopers() ([]*github.User, error) {
	appDevs := make(map[int64]*github.User)
	engineers, _, err := s.ghClient.Teams.ListTeamMembersByID(context.Background(), s.config.HipcampOrgID, s.config.EngineeringTeamID, nil)
	if err != nil {
		return nil, err
	}
	for _, engineer := range engineers {
		appDevs[*engineer.ID] = engineer
	}
	platformDevs, _, err := s.ghClient.Teams.ListTeamMembersByID(context.Background(), s.config.HipcampOrgID, s.config.PlatformTeamID, nil)
	if err != nil {
		return nil, err
	}
	for _, platformDev := range platformDevs {
		delete(appDevs, *platformDev.ID)
	}
	ret := make([]*github.User, len(appDevs))
	i := 0
	for _, appDev := range appDevs {
		ret[i] = appDev
		i++
	}
	return ret, nil
}

func (s *developerVelocity) numCommitsToMain(user *github.User, dayBeforeYesterday, yesterday string) (int, error) {
	if user.Name == nil || *user.Name == "" {
		return 0, fmt.Errorf("no username in %+v", user)
	}
	q := strings.ReplaceAll(query, "$username", *user.Name)
	q = strings.ReplaceAll(q, "$day_before_yesterday", dayBeforeYesterday)
	q = strings.ReplaceAll(q, "$yesterday", yesterday)
	commitSearchResult, _, err := s.ghClient.Search.Commits(context.Background(), q, nil)
	if err != nil {
		return 0, err
	}
	if commitSearchResult == nil || commitSearchResult.Total == nil {
		return 0, nil
	}
	return *commitSearchResult.Total, nil
}
