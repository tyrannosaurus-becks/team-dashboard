package metric_providers

import (
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers/asana"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers/github"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers/local_csv"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func All(config *models.Config) []models.MetricProvider {
	return []models.MetricProvider{
		asana.NewMetricProvider(config),
		github.NewMetricProvider(config),
		local_csv.NewMetricProvider(config.PathsToLocalCSVs),
	}
}
