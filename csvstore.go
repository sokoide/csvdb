package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type Record map[string]string // filele name -> value

type CSVStore struct {
	headers []string
	records []Record
	indices map[string]map[string]int // field name -> (value -> record index)
}

func NewCSVStoreFromReader(r io.Reader) (*CSVStore, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 1 {
		return nil, fmt.Errorf("empty CSV")
	}

	headers := make([]string, len(rows[0]))
	for i, h := range rows[0] {
		headers[i] = strings.TrimSpace(h)
	}
	indices := make(map[string]map[string]int)
	for _, h := range headers {
		indices[h] = make(map[string]int)
	}

	records := make([]Record, 0, len(rows)-1)
	for _, row := range rows[1:] {
		rec := make(Record)
		for i, val := range row {
			if i < len(headers) {
				rec[headers[i]] = strings.TrimSpace(val)
			}
		}
		records = append(records, rec)
		recordIdx := len(records) - 1
		for _, h := range headers {
			indices[h][rec[h]] = recordIdx
		}
	}

	return &CSVStore{headers, records, indices}, nil
}

func (s *CSVStore) Lookup(field string, value string) (Record, bool) {
	idx, ok := s.indices[field]
	if !ok {
		return nil, false
	}
	recordIdx, found := idx[value]
	if !found {
		return nil, false
	}
	return s.records[recordIdx], true
}
