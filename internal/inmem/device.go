package inmem

import "github.com/fanie42/dvras"

// Device TODO
type Device struct {
    id       dvras.DeviceID
    state    dvras.State
    sequence uint64
}
