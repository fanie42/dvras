package memory

import (
    "github.com/fanie42/dvras"
)

type cache struct {
    devices map[dvras.DeviceID]*dvras.Device
}

// NewCache TODO
func NewCache(
    repo dvras.Repository,
) dvras.Repository {
    return &cache{
        device: device,
    }
}

// Load TODO
func (c *cache) Load(id dvras.DeviceID) (*dvras.Device, error) {
    device, ok := c.devices[id]
    if !ok {
        return c.repo.Load(id)
    }

    return device, nil
}

// Save TODO
func (c *cache) Save(device *dvras.Device) error {
    err := c.repo.Save(device)

    repo.device = device

    return nil
}
