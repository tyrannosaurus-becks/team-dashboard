package dashboards

import (
	"testing"
)

func TestTimeConversion(t *testing.T) {
	testTime := float64(1575317847)
	asTime := fromDatadogTime(testTime)
	asPosix := toDatadogTime(asTime)
	if testTime != asPosix {
		t.Fatal("unexpected time")
	}
}
