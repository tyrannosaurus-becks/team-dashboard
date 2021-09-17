package local_csv

import (
	"errors"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/metric_providers/local_csv/metrics"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func NewMetricProvider(pathsToLocalCSVs map[models.CSVType]string) *MetricProvider {
	return &MetricProvider{
		pathsToLocalCSVs: pathsToLocalCSVs,
	}
}

type MetricProvider struct {
	pathsToLocalCSVs map[models.CSVType]string
}

func (m *MetricProvider) Calculate() ([]*models.Metric, error) {
	var ret []*models.Metric
	for csvType, path := range m.pathsToLocalCSVs {
		if csvType != models.GoogleForms {
			return nil, errors.New("only google forms is supported")
		}
		parser := metrics.NewGoogleFormParser(path)
		results, err := parser.Calculate()
		if err != nil {
			return nil, err
		}
		ret = append(ret, results...)
	}
	return ret, nil
}
