package queue

import (
	"fmt"
	"go-bert/event"
	"go-bert/handle"
	"log"
	"os"
	"time"
)

// Queue to represent all events to occur
type Queue struct {
	events        []event.Event
	maxTime       time.Duration
	maxIterations uint64
}

// RunEvents method to handle expected slice of events
func (q Queue) RunEvents() {
	fmt.Println(fmt.Sprintf("Running queue: %v events to process, Max operating time: %v\nLet's get this party started!", len(q.events), q.maxTime))

	timer := time.AfterFunc(q.maxTime, func() {
		log.Println("Program timer expired, exiting.")
		os.Exit(0)
	})
	defer timer.Stop()

	for i, event := range q.events {
		success, err := event.Run(i)
		handle.Error(err)
		fmt.Printf("Success on event %v: %v\n", i, success)

	}
}
