package models

type Config struct {
	// Datadog.
	DatadogClientAPIKey string `envconfig:"datadog_client_api_key"`

	// Github.
	GithubAccessToken string `envconfig:"github_access_token"`
	HipcampOrgID      int64  `envconfig:"hipcamp_org_id"`
	EngineeringTeamID int64  `envconfig:"engineering_team_id"`
	PlatformTeamID    int64  `envconfig:"platform_team_id"`

	// Asana.
	AsanaPersonalAccessToken string `envconfig:"asana_personal_access_token"`
	AsanaWorkspaceGid        string `envconfig:"asana_workspace_gid"`
	AsanaPlatformTeamGid     string `envconfig:"asana_platform_team_gid"`
	AsanaTypeFieldGid        string `envconfig:"asana_type_field_gid"`
	AsanaPriorityFieldGid    string `envconvig:"asana_priority_field_gid"`
}
