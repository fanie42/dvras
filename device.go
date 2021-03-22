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
    id       DeviceID
    state    State
    sequence uint64
    // changes  []Event
}

// NewDevice TODO
func NewDevice() *Device {
    return &Device{
        id:       DeviceID(uuid.New()),
        state:    Off,
        sequence: 0,
        // changes:  make([]Event, 0),
    }
}

// Apply TODO
func (device *Device) Apply(event Event) {
    switch event.(type) {
    case *StartedEvent:
        device.state = On
    case *StoppedEvent:
        device.state = Off
    case *DatapointAcquiredEvent:
    default:
        fmt.Println("unsupported event")
        return
    }
    device.sequence++
}

// ID TODO
func (device *Device) ID() DeviceID {
    return device.id
}

// State TODO
func (device *Device) State() State {
    return device.state
}

// Sequence TODO
func (device *Device) Sequence() uint64 {
    return device.sequence
}

// Changes TODO
// func (device *Device) Changes() []Event {
//     return device.changes
// }

// Start TODO
func (device *Device) Start(
    annotation string,
) error {
    device.Apply(
        &StartedEvent{
            ID: uuid.UUID(device.id),
            // Sequence:   device.sequence,
            Time:       time.Now(),
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
            ID: uuid.UUID(device.id),
            // Sequence:   device.Sequence(),
            Time:       time.Now(),
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
    if device.state == Off {
        err := device.Start("automatic start")
        if err != nil {
            return err
        }
    }
    device.Apply(
        &DatapointAcquiredEvent{
            ID: uuid.UUID(device.id),
            // Sequence: device.Sequence(),
            Time:     timestamp,
            Channel1: ch1,
            Channel2: ch2,
            PPS:      pps,
        },
    )

    return nil
}
