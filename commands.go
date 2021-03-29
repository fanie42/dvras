package dvras

import "time"

// Command TODO
type Command interface {
    Execute() ([]*Event, error)
}

// StartCommand TODO
type startCommand struct {
    device     *Device   `json:"device"`
    time       time.Time `json:"time"`
    annotation string    `json:"annotation"`
}

// Execute TODO
func (command *startCommand) Execute() ([]*Event, error) {
    events := []*Event{}

    events = append(
        events,
        &StartedEvent{
            id:         device.ID,
            time:       command.time,
            Annotation: command.annotation,
        },
    )

    return events, nil
}

// StartCommand TODO
type stopCommand struct {
    device     *Device   `json:"device"`
    time       time.Time `json:"time"`
    annotation string    `json:"annotation"`
}

// Execute TODO
func (command *stopCommand) Execute() ([]*Event, error) {
    events := []*Event{}

    events = append(
        events,
        &StartedEvent{
            id:         device.ID,
            time:       time.Now(),
            Annotation: command.annotation,
        },
    )

    return events, nil
}
