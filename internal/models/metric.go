package models

type Metric interface {
	Collect() (float64, error)
}
