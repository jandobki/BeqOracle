package oracle

import (
	"context"
	"fmt"

	"github.com/jandobki/beqoracle/server/internal/event"
)

type Service struct {
	store event.Store
}

func NewService() *Service {
	return &Service{
		store: &event.MemoryStore{},
	}
}

func (s *Service) CreateAnswer(ctx context.Context, key, value string) error {
	ev, err := s.store.GetLastEventByKey(ctx, key)
	if err != nil {
		return fmt.Errorf("can't retrieve answer: %v", err)
	}

	if ev.Event != event.EventTypeEmpty && ev.Event != event.EventTypeDelete {
		return fmt.Errorf("already exists")
	}

	err = s.store.AddEvent(ctx, event.Event{
		Event: event.EventTypeCreate,
		Key:   key,
		Data:  value,
	})

	if err != nil {
		return fmt.Errorf("can't create answer: %v", err)
	}

	return nil
}

func (s *Service) UpdateAnswer(ctx context.Context, key, value string) error {
	ev, err := s.store.GetLastEventByKey(ctx, key)
	if err != nil {
		return fmt.Errorf("can't update answer: %v", err)
	}

	if ev.Event != event.EventTypeCreate && ev.Event != event.EventTypeUpdate {
		return fmt.Errorf("doesn't exist")
	}

	err = s.store.AddEvent(ctx, event.Event{
		Event: event.EventTypeUpdate,
		Key:   key,
		Data:  value,
	})

	if err != nil {
		return fmt.Errorf("can't update answer: %v", err)
	}

	return nil
}

func (s *Service) GetAnswer(ctx context.Context, key string) (string, error) {
	ev, err := s.store.GetLastEventByKey(ctx, key)
	if err != nil {
		return "", fmt.Errorf("can't retrieve answer: %v", err)
	}

	if ev.Event != event.EventTypeCreate && ev.Event != event.EventTypeUpdate {
		return "", fmt.Errorf("doesn't exist")
	}

	return ev.Data, nil
}

func (s *Service) DeleteAnswer(ctx context.Context, key string) error {
	ev, err := s.store.GetLastEventByKey(ctx, key)
	if err != nil {
		return fmt.Errorf("can't retrieve answer: %v", err)
	}

	if ev.Event != event.EventTypeCreate && ev.Event != event.EventTypeUpdate {
		return fmt.Errorf("doesn't exist")
	}

	err = s.store.AddEvent(ctx, event.Event{
		Event: event.EventTypeDelete,
		Key:   key,
		Data:  "",
	})

	if err != nil {
		return fmt.Errorf("can't delete answer: %v", err)
	}

	return nil
}

type Event struct {
	Event string
	Value string
}

func (s *Service) GetAnswerHistory(ctx context.Context, key string, from, count int) ([]Event, int, error) {
	stored, to, err := s.store.GetEventsByKey(ctx, key, from, count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't retrieve history: %v", err)
	}
	res := make([]Event, len(stored))
	for i, e := range stored {
		res[i] = Event{Event: string(e.Event), Value: e.Data}
	}
	return res, to, nil
}
