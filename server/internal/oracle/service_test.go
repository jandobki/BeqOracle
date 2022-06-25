package oracle

import (
	"context"
	"testing"

	"github.com/jandobki/beqoracle/server/internal/event"
)

func TestCreateEventEmpty(t *testing.T) {
	service := Service{
		store: &event.MemoryStore{},
	}

	err := service.CreateAnswer(context.Background(), "testkey", "testvalue")
	if err != nil {
		t.Fatalf("CreateAnswer(...) on empty store failed")
	}
}

func TestCreateAnswerExisting(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "testkey1"},
				{Event: event.EventTypeUpdate, Key: "testkey1"},
				{Event: event.EventTypeCreate, Key: "testkey2"},
			},
		),
	}

	err := service.CreateAnswer(context.Background(), "testkey1", "testvalue")
	if err == nil {
		t.Fatalf("CreateAnswer(...\"testkey1\"...) succeeded while it existed")
	}

	err = service.CreateAnswer(context.Background(), "testkey2", "testvalue")
	if err == nil {
		t.Fatalf("CreateAnswer(...\"testkey2\"...) succeeded while it existed")
	}
}

func TestCreateAnswerDeleted(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "testkey"},
				{Event: event.EventTypeDelete, Key: "testkey"},
			},
		),
	}

	err := service.CreateAnswer(context.Background(), "testkey", "testvalue")
	if err != nil {
		t.Fatalf("CreateAnswer(...\"testkey\"...) failed while it was deleted")
	}
}

func TestUpdateAnswerEmpty(t *testing.T) {
	service := Service{
		store: &event.MemoryStore{},
	}

	err := service.UpdateAnswer(context.Background(), "testkey", "testvalue")
	if err == nil {
		t.Fatalf("UpdateAnswer(...) on empty store succeeded")
	}
}

func TestUpdateAnswerNotExisting(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "otherkey"},
			},
		),
	}

	err := service.UpdateAnswer(context.Background(), "testkey", "testvalue")
	if err == nil {
		t.Fatalf("UpdateAnswer(...\"testkey\"...) succeeded while it didn't exist")
	}
}

func TestGetAnswerEmpty(t *testing.T) {
	service := Service{
		store: &event.MemoryStore{},
	}

	a, err := service.GetAnswer(context.Background(), "testkey")
	if err == nil {
		t.Fatalf("UpdateAnswer(...) on empty store succeeded")
	}
	if a != "" {
		t.Fatalf("GetAnswer(...) on empty store returned value")
	}
}

func TestGetAnswerCreated(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "testkey", Data: "testvalue"},
			},
		),
	}

	a, err := service.GetAnswer(context.Background(), "testkey")
	if err != nil {
		t.Fatalf("GetAnswer(...\"testkey\"...) failed while it existed")
	}
	if a != "testvalue" {
		t.Fatalf("GetAnswer(...) returned wrong answer")
	}
}

func TestGetAnswerUpdated(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "testkey", Data: "testvalue1"},
				{Event: event.EventTypeUpdate, Key: "testkey", Data: "testvalue2"},
			},
		),
	}

	a, err := service.GetAnswer(context.Background(), "testkey")
	if err != nil {
		t.Fatalf("GetAnswer(...\"testkey\"...) failed while it existed")
	}
	if a != "testvalue2" {
		t.Fatalf("GetAnswer(...) returned wrong answer")
	}
}

func TestGetAnswerDeleted(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "testkey", Data: "testvalue1"},
				{Event: event.EventTypeDelete, Key: "testkey"},
			},
		),
	}

	a, err := service.GetAnswer(context.Background(), "testkey")
	if err == nil {
		t.Fatalf("GetAnswer(...\"testkey\"...) succeeded while it was deleted")
	}
	if a != "" {
		t.Fatalf("GetAnswer(...) returned answer while it was deleted")
	}
}

func TestGetAnswerHistoryAll(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "testkey", Data: "testvalue1"},
				{Event: event.EventTypeDelete, Key: "testkey"},
			},
		),
	}

	e, to, err := service.GetAnswerHistory(context.Background(), "testkey", 0, 10)
	if err != nil {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) failed")
	}
	if len(e) != 2 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) returned wrong number of events")
	}
	if to != 0 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) returned paging token while returning all")
	}
	if e[0].Event != string(event.EventTypeCreate) ||
		e[0].Value != "testvalue1" ||
		e[1].Event != string(event.EventTypeDelete) {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) returned wrong data")
	}
}

func TestGetAnswerHistoryPaged(t *testing.T) {
	service := Service{
		store: event.NewInitializedMemoryStore(
			[]event.Event{
				{Event: event.EventTypeCreate, Key: "testkey", Data: "testvalue1"},
				{Event: event.EventTypeUpdate, Key: "testkey", Data: "testvalue2"},
				{Event: event.EventTypeUpdate, Key: "testkey", Data: "testvalue3"},
				{Event: event.EventTypeUpdate, Key: "testkey", Data: "testvalue4"},
				{Event: event.EventTypeUpdate, Key: "testkey", Data: "testvalue5"},
				{Event: event.EventTypeDelete, Key: "testkey"},
			},
		),
	}

	e, to, err := service.GetAnswerHistory(context.Background(), "testkey", 0, 2)
	if err != nil {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) failed")
	}
	if len(e) != 2 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) returned wrong number of events")
	}
	if to == 0 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) returned zero paging token while returning part")
	}
	if e[0].Event != string(event.EventTypeCreate) ||
		e[0].Value != "testvalue1" {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) returned wrong data")
	}

	e, to, err = service.GetAnswerHistory(context.Background(), "testkey", to, 2)
	if err != nil {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 2nd page failed")
	}
	if len(e) != 2 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 2nd page returned wrong number of events")
	}
	if to == 0 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 2nd page returned zero paging token while returning part")
	}
	if e[0].Event != string(event.EventTypeUpdate) ||
		e[0].Value != "testvalue3" {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 2nd page returned wrong data")
	}

	e, to, err = service.GetAnswerHistory(context.Background(), "testkey", to, 2)
	if err != nil {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 3rd page failed")
	}
	if len(e) != 2 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 3rd page returned wrong number of events")
	}
	if to != 0 {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 3rd page returned paging token while it's the last page")
	}
	if e[0].Event != string(event.EventTypeUpdate) ||
		e[0].Value != "testvalue5" {
		t.Fatalf("GetAnswerHistory(...\"testkey\"...) 3rd page returned wrong data")
	}
}
