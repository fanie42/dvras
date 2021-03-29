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

// Status TODO
type Status int

const (
    // On TODO
    On Status = iota
    // Off TODO
    Off
)

// String TODO
func (status Status) String() string {
    switch status {
    case On:
        return "on"
    case Off:
        return "off"
    }

    return ""
}

// ParseStatus TODO
func ParseStatus(s string) Status {
    switch s {
    case "on":
        return On
    case "off":
        return Off
    }

    return -1
}

// Device TODO
type Device struct {
    id      deviceID
    status  Status
    version uint64
    // changes []Event
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

// Start TODO
func (device *Device) Start(
    start func() error,
    annotation string,
) Command {
    command := &startCommand{
        device:     device,
        time:       time.Now(),
        annotation: annotation,
    }

    return command
}

// Stop TODO
func (device *Device) Stop(
    annotation string,
) Command {
    command := &startCommand{
        device:     device,
        time:       time.Now(),
        annotation: annotation,
    }

    return command
}

// AcquireData TODO
func (device *Device) AcquireData(
    ew []int16,
    ns []int16,
    pps []int16,
) Command {
    // This should not be a command - it already happened. It's and event.
    command := &acquireDataCommand{
        device: device,
        time:   time.Now(),
        ew:     ew,
        ns:     ns,
        pps:    pps,
    }

    return command
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
// func (device *Device) Empty() {
//     device.changes = []Event{}
// }

// // Changes TODO
// func (device *Device) Changes() []Event {
//     return device.changes
// }
