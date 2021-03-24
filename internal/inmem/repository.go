package inmem

import (
    "fmt"
    "sync"

    "github.com/fanie42/dvras"
)

type gateway struct {
    sync.RWMutex
    events  []dvras.Event
    devices map[dvras.DeviceID]*Device
}

// New TODO
func New() dvras.Gateway {
    return &gateway{
        events:  make([]dvras.Event, 0),
        devices: make(map[dvras.DeviceID]*Device),
    }
}

// Load TODO
func (gw *gateway) Load(
    id dvras.DeviceID,
) (*dvras.Device, error) {
    gw.RLock()
    defer gw.RUnlock()

    device, ok := gw.devices[id]
    if !ok {
        newDevice := dvras.New(id)
        device = &Device{
            id:       newDevice.ID,
            state:    newDevice.State,
            sequence: newDevice.Sequence,
        }
        gw.devices[id] = device
    }

    return &dvras.Device{
        ID:       device.id,
        State:    device.state,
        Sequence: device.sequence,
    }, nil
}

// Save TODO
func (gw *gateway) Save(
    device *dvras.Device,
) error {
    gw.Lock()
    defer gw.Unlock()

    events := device.Changes()

    oldDevice, ok := gw.devices[device.ID]
    if !ok {
        newDevice := &Device{
            id:       device.ID,
            state:    device.State,
            sequence: device.Sequence,
        }
        gw.devices[device.ID] = newDevice
        oldDevice = newDevice
    }

    if device.Sequence != oldDevice.sequence+uint64(len(events)) {
        fmt.Printf(
            "old: %v ,,, new: %v ,,, events: %d\n",
            device.Sequence,
            oldDevice.sequence,
            len(events),
        )
        return dvras.SequenceConflictError{
            Have: device.Sequence,
            Want: oldDevice.sequence + uint64(len(events)),
        }
    }

    gw.events = append(gw.events, events...)
    gw.devices[device.ID].state = device.State
    gw.devices[device.ID].sequence = device.Sequence

    device.Empty()

    return nil
}
