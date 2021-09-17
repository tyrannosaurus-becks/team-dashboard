package local_csv

import (
	"fmt"
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
		models.GoogleForms: parts[0] + flag + "/internal/metric_providers/local_csv/fixtures/velocity.csv",
	}
	mp := NewMetricProvider(m)
	metrics, err := mp.Calculate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(metrics)
}
