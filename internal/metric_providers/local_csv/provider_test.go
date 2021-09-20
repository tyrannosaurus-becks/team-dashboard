package local_csv

import (
	"os"
	"strings"
	"testing"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func TestCalculate(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	flag := "team-dashboard"
	parts := strings.Split(wd, flag)
	m := map[models.CSVType]string{
		models.GoogleForms: parts[0] + flag + "/internal/metric_providers/local_csv/fixtures/velocity-05-26-21.csv",
	}
	mp := NewMetricProvider(m)
	metrics, err := mp.Calculate()
	if err != nil {
		t.Fatal(err)
	}
	if len(metrics) != 7 {
		t.Fatalf("expected %d, received %d", 7, len(metrics))
	}
}
