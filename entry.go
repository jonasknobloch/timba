package main

import (
	"strings"
	"time"
)

type entry struct {
	project     string
	description string
	start       time.Time
	end         time.Time
	duration    time.Duration
	tags        []string
}

func NewEntry(record []string) (*entry, error) {
	e := &entry{}

	e.project = record[3]
	e.description = record[5]

	layout := "2006-01-02 15:04:05"

	if t, err := time.Parse(layout, record[7]+" "+record[8]); err != nil {
		return nil, err
	} else {
		e.start = t
	}

	if t, err := time.Parse(layout, record[9]+" "+record[10]); err != nil {
		return nil, err
	} else {
		e.end = t
	}

	e.duration = e.end.Sub(e.start)

	if len(record[12]) > 0 {
		e.tags = strings.Split(record[12], ", ")
	}

	return e, nil
}
