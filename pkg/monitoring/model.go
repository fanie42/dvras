package monitoring

import (
    "time"

    "github.com/google/uuid"
)

// DeviceID TODO
type DeviceID uuid.UUID

// State TODO
type State int

const (
    // On TODO
    On State = iota
    // Off TODO
    Off
)

// String TODO
func (state State) String() string {
    switch state {
    case On:
        return "on"
    case Off:
        return "off"
    }

    return ""
}

// Device TODO
type Device struct {
    id    DeviceID
    state State
}

// DataPoint TODO
type DataPoint struct {
    deviceID  DeviceID
    timestamp time.Time
    ch1       []int16
    ch2       []int16
    pps       []int16
}
