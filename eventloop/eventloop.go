package eventloop

import (
    _ "fmt"
)

// Event represents a generic event.
type Event struct {
    Name string
    Data interface{}
}

// EventHandler defines the signature for event handling functions.
type EventHandler func(Event)

// EventEmitter represents an entity that can emit events.
type EventEmitter struct {
    handlers map[string][]EventHandler
}

// NewEventEmitter creates a new EventEmitter.
func NewEventEmitter() *EventEmitter {
    return &EventEmitter{
        handlers: make(map[string][]EventHandler),
    }
}

// On registers an event handler for the given event name.
func (emitter *EventEmitter) On(eventName string, handler EventHandler) {
    if _, ok := emitter.handlers[eventName]; !ok {
        emitter.handlers[eventName] = make([]EventHandler, 0)
    }
    emitter.handlers[eventName] = append(emitter.handlers[eventName], handler)
}

// Emit dispatches an event with the given name and data.
func (emitter *EventEmitter) Emit(eventName string, data interface{}) {
    event := Event{Name: eventName, Data: data}
    if handlers, ok := emitter.handlers[eventName]; ok {
        for _, handler := range handlers {
            handler(event)
        }
    }
}