package dashboards

import (
	"context"
	"time"

	dd "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func newDatadog(config *models.Config) *datadog {
	ctx := context.WithValue(
		context.Background(),
		dd.ContextAPIKeys,
		map[string]dd.APIKey{
			"apiKeyAuth": {
				Key: config.DatadogClientAPIKey,
			},
		},
	)
	configuration := dd.NewConfiguration()
	apiClient := dd.NewAPIClient(configuration)
	return &datadog{
		ctx:       ctx,
		apiClient: apiClient,
	}
}

type datadog struct {
	ctx       context.Context
	apiClient *dd.APIClient
}

func (d *datadog) Send(metrics []models.Metric) error {
	now := time.Now().UTC()
	for _, metric := range metrics {
		value, err := metric.Value()
		if err != nil {
			return err
		}
		if _, _, err := d.apiClient.MetricsApi.SubmitMetrics(d.ctx, *dd.NewMetricsPayload(
			[]dd.Series{*dd.NewSeries("platform.dashboard."+metric.Name(), [][]float64{
				{toDatadogTime(now), value},
			})}),
		); err != nil {
			return err
		}
	}
	return nil
}

// Timestamps should be in POSIX time in seconds,
// and cannot be more than ten minutes in the future
// or more than one hour in the past.
func toDatadogTime(t time.Time) float64 {
	return float64(t.Unix())
}

func fromDatadogTime(t float64) time.Time {
	return time.Unix(int64(t), 0)
}
