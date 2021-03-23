package inmem

import (
    "fmt"
    "sync"

    "github.com/fanie42/dvras"
)

type repository struct {
    sync.RWMutex
    events  []dvras.Event
    devices map[dvras.DeviceID]*Device
}

// New TODO
func New() dvras.Repository {
    return &repository{
        events:  make([]dvras.Event, 0),
        devices: make(map[dvras.DeviceID]*Device),
    }
}

// Load TODO
func (repo *repository) Load(
    id dvras.DeviceID,
) (*dvras.Device, error) {
    repo.RLock()
    defer repo.RUnlock()

    device, ok := repo.devices[id]
    if !ok {
        newDevice := dvras.NewDevice(id)
        device = &Device{
            id:       newDevice.ID,
            state:    newDevice.State,
            sequence: newDevice.Sequence,
        }
    }

    return &dvras.Device{
        ID:       device.id,
        State:    device.state,
        Sequence: device.sequence,
    }, nil
}

// Save TODO
func (repo *repository) Save(
    device *dvras.Device,
) error {
    repo.Lock()
    defer repo.Unlock()

    events := device.Changes()

    oldDevice, ok := repo.devices[device.ID]
    if !ok {
        newDevice := &Device{
            id:       device.ID,
            state:    device.State,
            sequence: device.Sequence,
        }
        repo.devices[device.ID] = newDevice
        oldDevice = newDevice
    }

    fmt.Printf("old: %v, new: %v\n", oldDevice.sequence, device.Sequence)

    for _, event := range events {
        repo.events = append(repo.events, events...)
        repo.devices[device.ID].state = device.State
        repo.devices[device.ID].sequence = device.Sequence
    }

    if device.Sequence == oldDevice.sequence+uint64(len(events)) {
    } else {
        return dvras.SequenceConflictError{
            Have: device.Sequence,
            Want: oldDevice.sequence + uint64(len(events)),
        }
    }

    fmt.Printf("old: %v, new: %v\n", oldDevice.sequence, device.Sequence)

    return nil
}
