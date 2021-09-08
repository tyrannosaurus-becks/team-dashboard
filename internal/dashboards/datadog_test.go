package dashboards

import (
	"fmt"
	"testing"
)

func TestTimeConversion(t *testing.T) {
	testTime := float64(1575317847)
	asTime := fromDatadogTime(testTime)
	fmt.Println(asTime.String())
	asPosix := toDatadogTime(asTime)
	if testTime != asPosix {
		t.Fatal("unexpected time")
	}
}
