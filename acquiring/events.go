package acquiring

import "time"

// EventType TODO
type EventType string

const (
    // SessionStarted TODO
    SessionStarted = EventType("SessionStarted")
    // SessionStopped TODO
    SessionStopped = EventType("SessionStopped")
    // DatumAcquired TODO
    DatumAcquired = EventType("DatumAcquired")
)

// Event TODO
type Event interface {
    Type() EventType
}

// SessionStartedEvent TODO
type SessionStartedEvent struct {
    Time time.Time `json:"time"`
}

// Type TODO
func (event *SessionStartedEvent) Type() EventType {
    return SessionStarted
}

// SessionStartedEventHandler TODO
// type SessionStartedEventHandler interface {
//     OnSessionStarted(*SessionStartedEvent)
// }

// SessionStoppedEvent TODO
type SessionStoppedEvent struct {
    Time time.Time `json:"time"`
}

// Type TODO
func (event *SessionStoppedEvent) Type() EventType {
    return SessionStopped
}

// SessionStoppedEventHandler TODO
// type SessionStoppedEventHandler interface {
//     OnSessionStopped(*SessionStoppedEvent)
// }

// DatumAcquiredEvent TODO
type DatumAcquiredEvent struct {
    Time    time.Time `json:"time"`
    Samples []int16   `json:"samples"`
}

// Type TODO
func (event *DatumAcquiredEvent) Type() EventType {
    return DatumAcquired
}

// DatumAcquiredEventHandler TODO
// type DatumAcquiredEventHandler interface {
//     OnDatumAcquired(*DatumAcquiredEvent)
// }
