package dvras

import (
    "fmt"
    "time"

    "github.com/google/uuid"
)

// DeviceID TODO
type DeviceID uuid.UUID

// String TODO
func (id DeviceID) String() string {
    return uuid.UUID(id).String()
}

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

// New TODO -
func New(id DeviceID) *Device {
    return &Device{
        ID:       id,
        State:    Off,
        Sequence: 0,
        changes:  make([]Event, 0),
    }
}

// raise TODO
func (device *Device) raise(event Event) {
    switch event.(type) {
    case *StartedEvent:
        device.State = On
    case *StoppedEvent:
        device.State = Off
    case *DatapointAcquiredEvent:
    default:
        fmt.Println("unsupported event")
        return
    }
    device.changes = append(device.changes, event)
    device.Sequence++
}

// Empty TODO
func (device *Device) Empty() {
    device.changes = []Event{}
}

// Changes TODO
func (device *Device) Changes() []Event {
    return device.changes
}

// Start TODO
func (device *Device) Start(
    annotation string,
) error {
    if device.State != Off {
        return fmt.Errorf("unable to start device, device already running")
    }

    device.raise(
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
    if device.State != On {
        return fmt.Errorf("unable to stop device, device not running")
    }

    device.raise(
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
    // if device.State != On {
    //     err := device.Start("automatic start")
    //     if err != nil {
    //         return err
    //     }
    // }

    fmt.Printf("%v\n", timestamp)

    device.raise(
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
