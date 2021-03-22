package acquisition

import (
    "time"

    "github.com/fanie42/dvras"
    "github.com/google/uuid"
)

// Device TODO
type Device struct {
    id    DeviceID
    state State

    events []dvras.Event
}

// NewDevice TODO
func NewDevice() *Device {
    return &Device{
        id:     DeviceID(uuid.New()),
        state:  Off,
        events: make([]dvras.Event, 0),
    }
}

// Apply TODO
func (device *Device) Apply(event dvras.Event) {
    switch event.(type) {
    case *dvras.StartedEvent:
        device.state = On
    case *dvras.StoppedEvent:
        device.state = Off
    case *dvras.DatapointAcquiredEvent:
    }
}

// ID TODO
func (device *Device) ID() DeviceID {
    return device.id
}

// State TODO
func (device *Device) State() State {
    return device.state
}

// Changes TODO
func (device *Device) Changes() []dvras.Event {
    return device.events
}

// Start TODO
func (device *Device) Start(
    annotation string,
) error {
    device.raise(&dvras.StartedEvent{
        DeviceID:   uuid.UUID(device.id),
        Annotation: annotation,
    })

    return nil
}

// Stop TODO
func (device *Device) Stop(
    annotation string,
) error {
    device.raise(&dvras.StoppedEvent{
        DeviceID:   uuid.UUID(device.id),
        Annotation: annotation,
    })

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
    device.raise(&dvras.DatapointAcquiredEvent{
        DeviceID:  uuid.UUID(device.id),
        Timestamp: timestamp,
        Channel1:  ch1,
        Channel2:  ch2,
        PPS:       pps,
    })

    return nil
}

func (device *Device) raise(event dvras.Event) {
    device.events = append(device.events, event)
    device.Apply(event)
}
