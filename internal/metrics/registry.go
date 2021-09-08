package metrics

import "github.com/tyrannosaurus-becks/team-dashboard/internal/models"

func All(config *models.Config) []models.Metric {
	return []models.Metric{
		newDeveloperVelocity(config),
		newPlatformStability(config),
		newHighRiskSecurityIssues(config),
	}
}
