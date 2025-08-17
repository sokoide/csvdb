package main

import (
	"strings"
	"testing"
)

func TestCSVStore(t *testing.T) {
	csvData := `hostname,cpus,memory
host1.com,4,8GB
host2.com,2,16GB
host3.com,8,32GB`

	r := strings.NewReader(csvData)
	store, err := NewCSVStoreFromReader(r)

	if err != nil {
		t.Fatalf("NewCSVStoreFromReader failed: %v", err)
	}

	// Successful lookup
	rec, found := store.Lookup("hostname", "host2.com")
	if !found {
		t.Errorf("Expected to find record with hostname host2.com, but did not")
	}
	if rec["cpus"] != "2" {
		t.Errorf("Expected cpus to be 2, but got %s", rec["cpus"])
	}
	if rec["memory"] != "16GB" {
		t.Errorf("Expected memory to be 16GB, but got %s", rec["memory"])
	}

	// Successful lookup
	rec, found = store.Lookup("cpus", "8")
	if !found {
		t.Errorf("Expected to find record with cpus 8, but did not")
	}
	if rec["hostname"] != "host3.com" {
		t.Errorf("Expected hostname to be host3.com, but got %s", rec["hostname"])
	}

	// Unsuccessful test (value not found)
	_, found = store.Lookup("hostname", "host4.com")
	if found {
		t.Errorf("Expected not to find record with hostname host4.com, but did")
	}

	// Unsuccessful lookup (field not found)
	_, found = store.Lookup("nonexistentfield", "somevalue")
	if found {
		t.Errorf("Expected not to find record with nonexistentfield, but did")
	}
}

func TestEmptyCSV(t *testing.T) {
	csvData := ``
	r := strings.NewReader(csvData)
	_, err := NewCSVStoreFromReader(r)
	if err == nil {
		t.Errorf("Expected an error for empty CSV, but got nil")
	}
}

func TestHeaderOnlyCSV(t *testing.T) {
	csvData := `hostname,cpus,memory`
	r := strings.NewReader(csvData)
	store, err := NewCSVStoreFromReader(r)
	if err != nil {
		t.Fatalf("NewCSVStoreFromReader failed for header-only CSV: %v", err)
	}

	_, found := store.Lookup("hostname", "host1.com")
	if found {
		t.Errorf("Expected not to find any record in a header-only CSV, but did")
	}
}
