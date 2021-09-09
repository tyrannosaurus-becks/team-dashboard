package models

type MetricType string

type Metric struct {
	Name  string
	Value float64
	Tags  *[]string
}
