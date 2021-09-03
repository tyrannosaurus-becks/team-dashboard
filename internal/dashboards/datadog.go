package dashboards

import (
	"context"

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
			"appKeyAuth": {
				Key: config.DatadogClientAppKey,
			},
		},
	)
	configuration := dd.NewConfiguration()
	apiClient := dd.NewAPIClient(configuration)
	return &datadog{
		ctx: ctx,
		apiClient: apiClient,
	}
}

type datadog struct {
	ctx context.Context
	apiClient *dd.APIClient
}

func (d *datadog) Send(metrics []models.Metric) error {
	for _, metric := range metrics {
		value, err := metric.Value()
		if err != nil {
			return err
		}
		if _, _, err := d.apiClient.MetricsApi.SubmitMetrics(d.ctx, dd.MetricsPayload{
			Series: []dd.Series{
				{
					Metric: "platform.dashboard." + metric.Name(),
					Points: [][]float64{
						{value},
					},
					Type: dd.PtrString(string(metric.Type())),
				},
			},
		}); err != nil {
			return err
		}
	}
	return nil
}