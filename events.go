package dvras

import (
    "time"

    "github.com/google/uuid"
)

// Event TODO
type Event interface {
}

// StartedEvent TODO
type StartedEvent struct {
    ID uuid.UUID `json:"id"`
    // Sequence   uint64    `json:"sequence"`
    Time       time.Time `json:"time"`
    Annotation string    `json:"annotation"`
}

// StoppedEvent TODO
type StoppedEvent struct {
    ID uuid.UUID `json:"id"`
    // Sequence   uint64    `json:"sequence"`
    Time       time.Time `json:"time"`
    Annotation string    `json:"annotation"`
}

// DatapointAcquiredEvent TODO
type DatapointAcquiredEvent struct {
    ID uuid.UUID `json:"id"`
    // Sequence uint64    `json:"sequence"`
    Time     time.Time `json:"time"`
    Channel1 []int16   `json:"channel_1"`
    Channel2 []int16   `json:"channel_2"`
    PPS      []int16   `json:"pps"`
}
