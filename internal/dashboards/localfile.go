package dashboards

import (
	"encoding/csv"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
	"os"
	"strconv"
	"time"
)

const fileName = "metrics.csv"

type localFile struct{}

func (f *localFile) Send(metrics []models.Metric) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	records := [][]string{}
	for _, metric := range metrics {
		value, err := metric.Value()
		if err != nil {
			return err
		}
		strValue := strconv.FormatFloat(value, 'f', 2, 64)
		tsValue := time.Now().UTC().Format(time.RFC3339)
		records = append(records, []string{
			strValue, tsValue,
		})
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}
