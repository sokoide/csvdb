# csvdb

A Go application that provides a simple CSV database-like interface for looking up records. It supports reading CSV data from both local files and HTTP URLs.

## Architecture

```mermaid
classDiagram
    class DataSource {
        <<interface>>
        +Lookup(field, value string) (Record, bool)
        +Refresh() error
    }
    
    class Record {
        <<type>>
        map[string]string
    }
    
    class CSVStore {
        -headers []string
        -indices map[string]map[string]Record
        +NewCSVStoreFromReader(r io.Reader) (*CSVStore, error)
        +Lookup(field, value string) (Record, bool)
    }
    
    class LocalCSV {
        -store *CSVStore
        +NewLocalCSV(path string) (*LocalCSV, error)
        +Lookup(field, value string) (Record, bool)
        +Refresh() error
    }
    
    class HTTPCSV {
        -store *CSVStore
        -url string
        +NewHTTPCSV(url string) (*HTTPCSV, error)
        +Lookup(field, value string) (Record, bool)
        +Refresh() error
    }
    
    DataSource <|.. LocalCSV : implements
    DataSource <|.. HTTPCSV : implements
    LocalCSV *-- CSVStore : aggregates
    HTTPCSV *-- CSVStore : aggregates
    CSVStore --> Record : uses
    DataSource --> Record : returns
```

## Usage

```go
// Local CSV file
ds, _ := NewLocalCSV("data.csv")
rec, found := ds.Lookup("hostname", "host1.com")

// HTTP CSV source
ds, _ := NewHTTPCSV("https://example.com/data.csv")
rec, found := ds.Lookup("cpus", "4")
```

## Build and Run

```bash
# Build the application
go build -o csvdb

# Run the application
go run .
```