package models

type MetricProvider interface {
	Calculate() ([]*Metric, error)
}
