package main

import (
    "log"

    "github.com/fanie42/dvras/internal/http/rest"
    "github.com/fanie42/dvras/internal/nats"
    "github.com/fanie42/dvras/internal/noop"
    "github.com/fanie42/dvras/internal/portaudio"
    pa "github.com/gordonklaus/portaudio"
    natsio "github.com/nats-io/nats.go"
)

func main() {
    connection, err := natsio.Connect("nats://172.18.30.100:4222")
    if err != nil {
        log.Fatalf("failed to connect to nats server: %v", err)
    }
    defer connection.Close()

    publisher := nats.NewPublisher(connection)
    repository := noop.NewRepository(publisher)

    err = pa.Initialize()
    if err != nil {
        log.Fatalf("failed to initialize portaudio: %v", err)
    }
    defer pa.Terminate()
    application := portaudio.New(
        &portaudio.Configuration{
            NSchannel:  0,
            EWchannel:  1,
            PPSchannel: 2,
        },
        repository,
    )
    defer application.Close()

    controller := rest.New(application)
    controller.Run()
}
