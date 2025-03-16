package bridge

import (
	"errors"
	"sync"
)

type Event struct {
	data *map[string]interface{}
}

type EventBus struct {
	events map[string]([]func(*Event) error)
	mut    sync.Mutex
}

func CreateEventBus() *EventBus {
	return &EventBus{
		events: make(map[string]([]func(*Event) error)),
	}
}

func (eb *EventBus) Publish(event_name string, event *Event) error {
	eb.mut.Lock()
	defer eb.mut.Unlock()
	subscribers, ok := eb.events[event_name]

	if ok {

		for _, subscriber := range subscribers {
			subscriber(event)
		}
		return nil
	}
	return errors.New("event doesnt exist")
}

func (eb *EventBus) Subscribe(event_name string, callback func(*Event) error) {
	eb.mut.Lock()
	defer eb.mut.Unlock()
	subscribers, exists := eb.events[event_name]

	if exists {
		eb.events[event_name] = append(subscribers, callback)
	} else {
		eb.events[event_name] = [](func(*Event) error){callback}
	}
}

func CreateEvent() *Event {
	return &Event{}
}
func CreateEventAction(action Action) *Event {
	new_event := &Event{}
	new_event.Add("packet", action)
	return new_event
}

func (event *Event) GetData() *map[string]interface{} {
	return event.data
}

func (event *Event) Add(key string, value interface{}) {
	(*event.data)[key] = value
}
