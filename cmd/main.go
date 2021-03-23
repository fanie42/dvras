package main

import (
    "fmt"

    "github.com/fanie42/dvras"
    "github.com/fanie42/dvras/internal/http/rest"
    "github.com/fanie42/dvras/internal/inmem"
    "github.com/fanie42/dvras/internal/portaudio"
    "github.com/google/uuid"
    pa "github.com/gordonklaus/portaudio"
)

func main() {
    // dbpool, err := pgxpool.Connect(
    //     context.Background(),
    //     "postgres://postgres:admin@172.18.30.100:5432/dvras",
    // )
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer dbpool.Close()
    // repository := timescaledb.New(dbpool)
    repository := inmem.New()

    err := pa.Initialize()
    if err != nil {
        fmt.Printf("failed to initialize portaudio: %v", err)
        return
    }
    defer pa.Terminate()
    app := portaudio.New(
        &portaudio.Config{
            SampleRate: 44100,
            DeviceID:   dvras.DeviceID(uuid.New()),
        },
        repository,
    )
    defer app.Close()

    controller := rest.New(app)
    controller.Run()
}
