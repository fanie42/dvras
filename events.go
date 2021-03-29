package dvras

import (
    "time"
)

// Event TODO
type Event interface {
    Apply()
}

// StartedEvent TODO
type startedEvent struct {
    device     *Device
    time       time.Time `json:"time"`
    annotation string    `json:"annotation"`
}

func (event *startedEvent) Apply() {
    event.device.state = On
}

// StoppedEvent TODO
type stoppedEvent struct {
    device     *Device
    time       time.Time `json:"time"`
    annotation string    `json:"annotation"`
}

func (event *stoppedEvent) Apply() {
    event.device.state = Off
}

// DataAcquiredEvent TODO
type dataAcquiredEvent struct {
    device *Device
    time   time.Time `json:"time"`
    ew     []int16   `json:"ew"`
    ns     []int16   `json:"ns"`
    pps    []int16   `json:"pps"`
}

func (event *dataAcquiredEvent) Apply() {
    event.device.data = &Data{
        ew:  event.device.ew,
        ns:  event.device.ns,
        pps: event.device.pps,
    }
}
