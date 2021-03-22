package inmem

import (
    "fmt"

    "github.com/fanie42/dvras"
    "github.com/fanie42/dvras/pkg/monitoring"
)

type repository struct {
    devices map[monitoring.DeviceID]*monitoring.Device
    events  map[monitoring.EventID]*dvras.Event
}

// New TODO
func New() monitoring.Repository {
    return &repository{
        devices: make(map[monitoring.DeviceID]*monitoring.Device),
        events:  make(map[monitoring.EventID]*dvras.Event),
    }
}

// Load TODO
func (repo *repository) Load(
    id monitoring.DeviceID,
) (*monitoring.Device, error) {
    device, ok := repo.devices[id]
    if !ok {
        return nil, fmt.Errorf("no device with id: %v", id)
    }

    return device, nil
}

// Save TODO
func (repo *repository) Save(
    device *monitoring.Device,
) error {
    for _, event := range device.Changes() {
        repo.events[event.ID()] = event
        event.Apply(device)
    }
    repo.devices[device.ID] = device

    return nil
}
