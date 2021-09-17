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
	anonymizedResponses := make(map[string][]int)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for columnIdx, cellValue := range record {
			if columnIdx == 0 && cellValue == "" {
				// For normal responses, there's always a timestamp. The last
				// row has no timestamp, but only aggregations. We don't want
				// to grab that one.
				break
			}
			// We're in the header row, build the header cellValue to its index.
			if rowIdx == 0 {
				columnIdxToColumnName[columnIdx] = cellValue
			} else {
				// We're in a content row and simply want to capture its value,
				// if it's quantitative.
				num, err := strconv.Atoi(cellValue)
				if err != nil {
					// It's not an int.
					continue
				}
				columnName := columnIdxToColumnName[columnIdx]
				values, ok := anonymizedResponses[columnName]
				if !ok {
					values = []int{}
				}
				anonymizedResponses[columnName] = append(values, num)
			}
		}
		rowIdx++
	}
	for fieldName, values := range anonymizedResponses {
		for _, value := range values {
			metrics = append(metrics, &models.Metric{
				Name:  fieldName,
				Value: float64(value), // TODO each one of these isn't a gauge. What is it?
			})
		}
	}
	return metrics, nil
}
