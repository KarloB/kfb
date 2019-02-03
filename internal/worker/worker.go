package worker

import (
	"log"

	"github.com/KarloB/kfb/internal/firebase"
)

// Worker sync worker
type Worker struct {
	firebase *firebase.Client
	stop     chan struct{}
}

// New init worker
func New(fc *firebase.Client, stop chan struct{}) *Worker {
	return &Worker{
		firebase: fc,
		stop:     stop,
	}
}

// Watch watch db
func (t *Worker) Watch() error {
	err := t.firebase.Watcher(t, nil)
	if err != nil {
		return err
	}
	return nil
}

// Do does something with event
func (t *Worker) Do(event string, data interface{}) error {
	log.Printf("Worker that does something with firebase data")
	log.Printf("Event: %s", event)
	log.Printf("Data: %+v", data)
	return nil
}
