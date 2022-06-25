package event

type EventType string

const (
	EventTypeEmpty  EventType = ""
	EventTypeCreate EventType = "create"
	EventTypeUpdate EventType = "update"
	EventTypeDelete EventType = "delete"
)

type Event struct {
	Event EventType
	Key   string
	Data  string
}
