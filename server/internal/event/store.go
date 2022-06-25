package event

import (
	"context"
)

type Store interface {
	AddEvent(ctx context.Context, event Event) error
	GetLastEventByKey(ctx context.Context, key string) (Event, error)
	GetEventsByKey(ctx context.Context, key string, from, count int) ([]Event, int, error)
}

type MemoryStore struct {
	events []Event
}

func NewInitializedMemoryStore(events []Event) *MemoryStore {
	return &MemoryStore{events: events}
}

func (s *MemoryStore) AddEvent(ctx context.Context, event Event) error {
	s.events = append(s.events, event)
	return nil
}

func (s *MemoryStore) GetLastEventByKey(ctx context.Context, key string) (Event, error) {
	for i := len(s.events) - 1; i >= 0; i-- {
		if s.events[i].Key == key {
			return s.events[i], nil
		}
	}
	return Event{}, nil
}

const memMaxPageSize = 100

func (s *MemoryStore) GetEventsByKey(ctx context.Context, key string, from, count int) ([]Event, int, error) {
	if count > memMaxPageSize {
		count = memMaxPageSize
	}
	res := make([]Event, 0, count)

	for i := from; i < len(s.events); i++ {
		if len(res) == count {
			return res, i, nil
		}

		if s.events[i].Key == key {
			res = append(res, s.events[i])
		}
	}

	return res, 0, nil
}
