package main

import (
	log "github.com/sirupsen/logrus"
)

func search(ds DataSource, field string, value string) {
	rec, found := ds.Lookup(field, value)
	if found {
		log.Infof("[%s,%s] Found %+v", field, value, rec)
	} else {
		log.Infof("[%s,%s] Not found", field, value)
	}
}

func main() {
	ds, _ := NewLocalCSV("data.csv")
	search(ds, "hostname", "host1.com")
	search(ds, "cpus", "2")
	search(ds, "cpus", "100")
}
