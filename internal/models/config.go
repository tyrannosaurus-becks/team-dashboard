package models

type Config struct {
	// Datadog.
	DatadogClientAPIKey string `envconfig:"datadog_client_api_key"`

	// Github.
	GithubAccessToken string `envconfig:"github_access_token"`
	HipcampOrgID      int64  `envconfig:"hipcamp_org_id"`
	EngineeringTeamID int64  `envconfig:"engineering_team_id"`
	PlatformTeamID    int64  `envconfig:"platform_team_id"`
}
