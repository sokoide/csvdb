package main

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Record map[string]string // filele name -> value

type CSVStore struct {
	headers []string
	indices map[string]map[string]Record // field name -> (value -> record)
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

	headers := rows[0]
	indices := make(map[string]map[string]Record)
	for _, h := range headers {
		indices[h] = make(map[string]Record)
	}

	for _, row := range rows[1:] {
		rec := make(Record)
		for i, val := range row {
			if i < len(headers) {
				rec[headers[i]] = val
			}
		}
		for _, h := range headers {
			indices[h][rec[h]] = rec
		}
	}

	return &CSVStore{headers, indices}, nil
}

func (s *CSVStore) Lookup(field string, value string) (Record, bool) {
	idx, ok := s.indices[field]
	if !ok {
		return nil, false
	}
	rec, found := idx[value]
	return rec, found
}
