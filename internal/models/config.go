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
	AsanaPriorityFieldGid    string `envconfig:"asana_priority_field_gid"`

	// CSVs.
	// The absolute path to any local CSVs to read as a comma-separated string.
	// Ex. "/Users/beccapetrin/go/team-dashboard.csv,/Users/beccapetrin.com.csv"
	PathsToLocalCSVs map[CSVType]string `envconfig:"paths_to_local_csvs"`
}
