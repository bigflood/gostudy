package cmd

import (
	"github.com/bigflood/gostudy/todo/store"
	"log"
	"time"
)

func NewLogging(next Store) Store {
	return loggingSvc{next}
}

type loggingSvc struct {
	next Store
}

func (svc loggingSvc) Add(desc string) (err error) {
	defer func(startTime time.Time) {
		took := time.Since(startTime)
		log.Println("Add", desc, "->", err, "took:", took)
	}(time.Now())

	err = svc.next.Add(desc)
	return
}

func (svc loggingSvc) List(filter store.Filter) (tasks []store.Task, err error) {
	defer func(startTime time.Time) {
		took := time.Since(startTime)
		log.Println("List", filter, "->", tasks, err, "took:", took)
	}(time.Now())

	tasks, err = svc.next.List(filter)
	return
}

func (svc loggingSvc) Done(index int) (err error) {
	defer func(startTime time.Time) {
		took := time.Since(startTime)
		log.Println("Done", index, "->", err, "took:", took)
	}(time.Now())

	err = svc.next.Done(index)
	return
}
