package models

type MetricType string

type Metric interface {
	Name() string
	Value() (float64, error)
}
