package main

import (
    "context"
    "fmt"
    "log"

    "github.com/fanie42/dvras/internal/acquisition/http/rest"
    "github.com/fanie42/dvras/internal/acquisition/portaudio"
    "github.com/fanie42/dvras/internal/acquisition/timescaledb"
    "github.com/fanie42/dvras/pkg/acquisition"
    "github.com/google/uuid"
    pa "github.com/gordonklaus/portaudio"
    "github.com/jackc/pgx/v4/pgxpool"
)

func main() {
    dbpool, err := pgxpool.Connect(
        context.Background(),
        "postgres://postgres:admin@172.18.30.100:5432/dvras",
    )
    if err != nil {
        log.Fatal(err)
    }
    defer dbpool.Close()
    repository := timescaledb.New(dbpool)

    err := pa.Initialize()
    if err != nil {
        fmt.Printf("failed to initialize portaudio: %v", err)
        return
    }
    defer pa.Terminate()
    app := portaudio.New(
        &portaudio.Config{
            SampleRate: 44100,
            DeviceID:   acquisition.DeviceID(uuid.New()),
        },
        repository,
    )
    defer app.Close()

    controller := rest.New(app)
    controller.Run()
}
