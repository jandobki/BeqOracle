package event

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PgStore struct {
	db *sql.DB
}

func getPsqlInfo() string {
	host := os.Getenv("POSTGRES_HOST")
	db := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	return fmt.Sprintf("host=%s port=5432 user=%s password=%s sslmode=disable database=%s", host, user, password, db)
}

const initQuery = `
CREATE TABLE IF NOT EXISTS events (
    "id" serial primary key,
    "event" varchar(20),
    "key" TEXT,
    "value" TEXT
)`

func InitDB(ctx context.Context, dbName string) error {
	log.Print("initializing database... ")

	db, err := sql.Open("postgres", getPsqlInfo())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, initQuery)

	return err
}

func NewPqStore() (*PgStore, error) {
	db, err := sql.Open("postgres", getPsqlInfo())

	if err != nil {
		return nil, err
	}

	return &PgStore{
		db: db,
	}, nil
}

const addQuery = `INSERT INTO public.events ("event", "key", "value") VALUES($1, $2, $3);`

func (s *PgStore) AddEvent(ctx context.Context, event Event) error {
	_, err := s.db.ExecContext(ctx, addQuery, event.Event, event.Key, event.Data)

	return err
}

const selectLastQuery = `SELECT "event", "value" FROM public.events WHERE "key" = $1 ORDER BY "id" DESC LIMIT 1;`

func (s *PgStore) GetLastEventByKey(ctx context.Context, key string) (Event, error) {
	var event, value string
	err := s.db.QueryRowContext(ctx, selectLastQuery, key).Scan(&event, &value)
	if err != nil && err != sql.ErrNoRows {
		return Event{}, err
	}

	return Event{Event: EventType(event), Key: key, Data: value}, nil
}

const selectAllQuery = `SELECT "id", "event", "value" FROM public.events WHERE "id" >= $2 AND "key" = $1 ORDER BY "id" LIMIT $3`
const pgMaxPageSize = 100

func (s *PgStore) GetEventsByKey(ctx context.Context, key string, from, count int) ([]Event, int, error) {
	if count > pgMaxPageSize {
		count = pgMaxPageSize
	}
	res := make([]Event, 0, count)

	rows, err := s.db.QueryContext(ctx, selectAllQuery, key, from, count+1)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	next_id := 0
	for i := 0; rows.Next(); i++ {
		var id int
		var event, value string
		rows.Scan(&id, &event, &value)
		if i >= count {
			next_id = id
			break
		}
		res = append(res, Event{
			Event: EventType(event),
			Key:   key,
			Data:  value})
	}

	return res, next_id, nil
}
