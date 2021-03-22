package portaudio

import (
    dvras "github.com/fanie42/dvras/pkg/acquisition"
)

// Config TODO
type Config struct {
    SampleRate int            `json:"sample_rate"`
    DeviceID   dvras.DeviceID `json:"device_id"`
}
