package main

import (
    "fmt"

    "github.com/fanie42/dvras"
    "github.com/fanie42/dvras/internal/http/rest"
    "github.com/fanie42/dvras/internal/memory"
    "github.com/fanie42/dvras/internal/portaudio"
    "github.com/google/uuid"
    pa "github.com/gordonklaus/portaudio"
)

func main() {
    repository := memory.New()

    err = pa.Initialize()
    if err != nil {
        fmt.Printf("failed to initialize portaudio: %v", err)
        return
    }
    defer pa.Terminate()
    application := portaudio.New(
        &portaudio.Config{
            SampleRate: 44100,
            DeviceID:   dvras.DeviceID(uuid.New()),
        },
        gateway,
    )
    defer application.Close()

    controller := rest.New(application)
    controller.Run()
}
