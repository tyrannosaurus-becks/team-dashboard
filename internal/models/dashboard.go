package models

type Dashboard interface {
	Send(metrics []Metric) error
}
