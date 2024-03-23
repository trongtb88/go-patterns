package main

import (
	"fmt"
	"sync"
	"time"
)

type Observer interface {
	NotifyCallback(Event)
}

type Event struct {
	Data int
}
type eventObserver struct {
	id   int
	time time.Time
}
type eventSubject struct {
	observers sync.Map
}

type Subject interface {
	AddListener(Observer)
	RemoveListener(Observer)
	Notify(Event)
}

func (e *eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Received: %d after %v\n", event.Data, time.Since(e.time))
}

func (s *eventSubject) AddListener(obs Observer) {
	s.observers.Store(obs, 1)
}

func (s *eventSubject) RemoveListener(obs Observer) {
	s.observers.Delete(obs)
}
func (s *eventSubject) Notify(e Event) {
	s.observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(e)
		return true
	})
}
func main() {
	n := eventSubject{
		observers: sync.Map{},
	}
	var observer1 = eventObserver{1, time.Now()}
	var observer2 = eventObserver{2, time.Now()}

	n.AddListener(&observer1)
	n.AddListener(&observer2)
	n.Notify(Event{
		Data: 5555,
	})

}
