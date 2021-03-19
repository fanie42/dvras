package inmem

import (
    "fmt"

    dvras "github.com/fanie42/dvras/pkg"
)

type repository struct {
    devices map[dvras.DeviceID]*dvras.Device
}

// New TODO
func New() dvras.DeviceRepository {
    repo := &repository{
        devices: make(map[dvras.DeviceID]*dvras.Device),
    }

    return repo
}

// Save TODO
func (repo *repository) Save(
    device *dvras.Device,
) error {
    repo.devices[device.ID()] = device

    return nil
}

// GetDeviceByID TODO
func (repo *repository) GetDeviceByID(
    id dvras.DeviceID,
) (*dvras.Device, error) {
    device, ok := repo.devices[id]
    if !ok {
        return nil, fmt.Errorf("no device with this id")
    }

    return device, nil
}
