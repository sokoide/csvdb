package main

import (
	"os"
)

type LocalCSV struct {
	store *CSVStore
}

func NewLocalCSV(path string) (*LocalCSV, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	store, err := NewCSVStoreFromReader(f)
	if err != nil {
		return nil, err
	}
	return &LocalCSV{store}, nil
}

func (l *LocalCSV) Lookup(field, value string) (Record, bool) {
	return l.store.Lookup(field, value)
}

func (h *LocalCSV) Refresh() error {
	// TODO:
	return nil
}
