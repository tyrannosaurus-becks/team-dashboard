package metrics

import "time"

var (
	iso8601        = "2006-01-02"
	lastSevenDays  = -7 * 24 * time.Hour
	lastThirtyDays = -3 * 24 * time.Hour
)
