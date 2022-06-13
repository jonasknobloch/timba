package main

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"github.com/jonasknobloch/timba/tim"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no import file specified")
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()

	backup := tim.NewBackup()

	tasks := make(map[string]*tim.Task)   // description -> task
	groups := make(map[string]*tim.Group) // project -> group

	getGroup := func(group string) *tim.Group {
		if group == "" {
			return nil
		}

		if g, ok := groups[group]; ok {
			return g
		}

		now := time.Now()

		g := tim.NewGroup(uuid.New().String(), group, now, now)

		backup.AddGroup(g)

		groups[group] = g

		return g
	}

	getTask := func(task, group string) *tim.Task {
		if t, ok := tasks[task]; ok {
			return t
		}

		now := time.Now()

		t := tim.NewTask(uuid.New().String(), task, now, now)
		g := getGroup(group)

		backup.AddTask(t, g)

		tasks[task] = t

		return t
	}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		entry, err := NewEntry(record)

		if err != nil {
			log.Fatal(err)
		}

		var t *tim.Task

		if entry.description == "" {
			t = getTask(entry.project, "")
		} else {
			t = getTask(entry.description, entry.project)
		}

		t.AddRecord(tim.NewRecord(entry.start, entry.end, strings.Join(entry.tags, ", ")))
	}

	if j, err := backup.ToJSON(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(j)
	}
}
