package dvras

import (
    "log"
    "time"
)

// Not data, it is the device. BUT the data is part of its current state. The
// repository needn't be event sourced.

// ID TODO
type ID time.Time

// Data TODO
type Data struct {
    id       ID
    deviceID device.ID
    ns       []int32
    ew       []int32
    pps      []int32
    changes  []Event
}

// New TODO
func New(id ID, ns, ew, pps []int32) *Data {
    return &Data{
        id:      id,
        ns:      ns,
        ew:      ew,
        pps:     pps,
        changes: []Event{},
    }
}

func (data *Data) apply(event Event) {
    switch e := event.(type) {
    case *AcquiredEvent:
        data.ns = e.NS
        data.ew = e.EW
        data.pps = e.PPS
    default:
        log.Printf("unsupported event type: %v", e)
    }
}

// Changes TODO
func (data *Data) Changes() []Event {
    return data.changes
}

// Empty TODO
func (data *Data) Empty() {
    data.changes = []Event{}
}

// Acquire TODO
func (data *Data) Acquire(deviceID device.ID) error {
    data.raise(
        &AcquiredEvent{
            ID:       data.id,
            DeviceID: deviceID,
            NS:       data.ns,
            EW:       data.ew,
            PPS:      data.pps,
        },
    )

    return nil
}

// Raise TODO
func (data *Data) raise(event Event) {
    data.changes = append(data.changes, event)
    data.apply(event)
}
