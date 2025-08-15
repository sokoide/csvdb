# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go application called `csvdb` that provides a simple CSV database-like interface for looking up records. It supports reading CSV data from both local files and HTTP URLs.

## Build and Run Commands

```bash
# Build the application
go build -o csvdb

# Run the application directly
go run .

# Run with Go modules support
go mod tidy
go run main.go

# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Run tests (if any are added)
go test ./...
```

## Architecture

The application follows a clean architecture pattern with:

- **DataSource interface** (`datasource.go`): Defines the contract for data sources with `Lookup()` and `Refresh()` methods
- **CSVStore** (`csvstore.go`): Core CSV parsing and indexing engine that creates indices for all fields
- **LocalCSV** (`localcsv.go`): Implementation for reading CSV files from local filesystem
- **HTTPCSV** (`httpcsv.go`): Implementation for fetching CSV data from HTTP URLs
- **Record type**: `map[string]string` representing a single CSV row

### Key Components

- `CSVStore.indices`: `map[string]map[string]Record` - Creates a nested map index for fast field-value lookups
- `NewCSVStoreFromReader()`: Factory function that parses CSV and builds indices for all columns
- Both LocalCSV and HTTPCSV wrap CSVStore and implement the DataSource interface

### Data Flow

1. CSV data is read through io.Reader interface
2. First row becomes headers, subsequent rows become records
3. Indices are built for every field to enable O(1) lookups
4. Lookup queries use field name and value to find matching records

## Dependencies

- `github.com/sirupsen/logrus`: Logging library
- Standard library packages: `encoding/csv`, `net/http`, `os`, `io`
- Test dependencies include `testify` (present in go.sum but no tests currently exist)

## Current Limitations

- `Refresh()` methods in both LocalCSV and HTTPCSV are not implemented (marked TODO)
- No tests currently exist despite testify being in dependencies
- No error handling for malformed CSV data beyond basic parsing errors