package portaudio

import (
    "github.com/fanie42/dvras"
)

// Config TODO
type Config struct {
    SampleRate int            `json:"sample_rate"`
    DeviceID   dvras.DeviceID `json:"device_id"`
}
