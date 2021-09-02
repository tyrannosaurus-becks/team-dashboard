package models

type MetricType string
const (
	Count MetricType = "count"
	Gauge MetricType = "gauge"
	Rate MetricType = "rate"
)

type Metric interface {
	Name() string
	Type() MetricType
	Value() (float64, error)
}
