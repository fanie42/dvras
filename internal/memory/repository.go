package memory

import (
    "sync"

    "github.com/fanie42/dvras"
)

type repository struct {
    sync.RWMutex
    devices map[dvras.DeviceID]*dvras.Device
}

// New TODO
func New(
    device *dvras.Device,
) dvras.Repository {
    return &repository{
        device: device,
    }
}

// Load TODO
func (repo *repository) Load() *dvras.Device {
    repo.RLock()
    defer repo.RUnlock()

    return repo.device
}

// Save TODO
func (repo *repository) Save(device *dvras.Device) error {
    repo.Lock()
    defer repo.Unlock()

    repo.device = device

    return nil
}
