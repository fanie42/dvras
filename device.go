package dvras

import (
    "fmt"
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
    ID       DeviceID
    State    State
    Sequence uint64
    changes  []Event
}

// NewDevice TODO
func NewDevice(id DeviceID) *Device {
    return &Device{
        ID:       id,
        State:    Off,
        Sequence: 0,
        changes:  make([]Event, 0),
    }
}

// Apply TODO
func (device *Device) Apply(event Event) {
    switch event.(type) {
    case *StartedEvent:
        device.State = On
        device.Sequence++
    case *StoppedEvent:
        device.State = Off
        device.Sequence++
    case *DatapointAcquiredEvent:
    default:
        fmt.Println("unsupported event")
        return
    }
    device.changes = append(device.changes, event)
}

// ID TODO
// func (device *Device) ID() DeviceID {
//     return device.id
// }

// // State TODO
// func (device *Device) State() State {
//     return device.state
// }

// // Sequence TODO
// func (device *Device) Sequence() uint64 {
//     return device.sequence
// }

// Changes TODO
func (device *Device) Changes() []Event {
    return device.changes
}

// Start TODO
func (device *Device) Start(
    annotation string,
) error {
    device.Apply(
        &StartedEvent{
            id: device.ID,
            // Sequence:   device.sequence,
            time:       time.Now(),
            Annotation: annotation,
        },
    )

    return nil
}

// Stop TODO
func (device *Device) Stop(
    annotation string,
) error {
    device.Apply(
        &StoppedEvent{
            id: device.ID,
            // Sequence:   device.Sequence(),
            time:       time.Now(),
            Annotation: annotation,
        },
    )

    return nil
}

// AcquireDataPoint TODO
func (device *Device) AcquireDataPoint(
    timestamp time.Time,
    ch1 []int16,
    ch2 []int16,
    pps []int16,
) error {
    if device.State == Off {
        err := device.Start("automatic start")
        if err != nil {
            return err
        }
    }
    device.Apply(
        &DatapointAcquiredEvent{
            id: device.ID,
            // Sequence: device.Sequence(),
            time:     timestamp,
            Channel1: ch1,
            Channel2: ch2,
            PPS:      pps,
        },
    )

    return nil
}
