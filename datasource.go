package main

type DataSource interface {
	Lookup(field, value string) (Record, bool)
	Refresh() error
}
