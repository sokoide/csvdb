package main

import (
	"fmt"
	"net/http"
)

type HTTPCSV struct {
	store *CSVStore
	url   string
}

func NewHTTPCSV(url string) (*HTTPCSV, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	store, err := NewCSVStoreFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return &HTTPCSV{store: store, url: url}, nil
}

func (h *HTTPCSV) Lookup(field, value string) (Record, bool) {
	return h.store.Lookup(field, value)
}

func (h *HTTPCSV) Refresh() error {
	// TODO:
	return nil
}
