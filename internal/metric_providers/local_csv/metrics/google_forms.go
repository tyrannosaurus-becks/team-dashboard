package metrics

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func NewGoogleFormParser(pathToLocalCSV string) *GoogleFormParser {
	return &GoogleFormParser{
		pathToLocalCSV: pathToLocalCSV,
	}
}

type GoogleFormParser struct {
	pathToLocalCSV string
}

func (f *GoogleFormParser) Calculate() ([]*models.Metric, error) {
	var metrics []*models.Metric

	csvFile, err := os.Open(f.pathToLocalCSV)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvFile)
	rowIdx := 0
	columnIdxToColumnName := make(map[int]string)
	anonymizedResponses := make(map[string][]float64)
	inFinalRow := false
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for columnIdx, cellValue := range record {
			switch {
			case rowIdx == 0:
				// We're in the header row, build the header cellValue to its index.
				columnIdxToColumnName[columnIdx] = cellValue
			case (columnIdx == 0 && cellValue == "") || inFinalRow:
				// For normal responses, there's always a timestamp. The last
				// row has no timestamp, but only aggregations. We're in the last
				// row.
				inFinalRow = true
				num, err := strconv.ParseFloat(cellValue, 64)
				if err != nil {
					// It's not a decimal.
					continue
				}
				columnName := columnIdxToColumnName[columnIdx]
				values, ok := anonymizedResponses[columnName]
				if !ok {
					values = []float64{}
				}
				anonymizedResponses[columnName] = append(values, num)
			default:
				// We're in one of the individual data rows. We only want the
				// aggregation at the end.
				continue
			}
		}
		rowIdx++
	}
	for fieldName, values := range anonymizedResponses {
		for _, value := range values {
			metrics = append(metrics, &models.Metric{
				Name:  fieldName,
				Value: value,
			})
		}
	}
	return metrics, nil
}
