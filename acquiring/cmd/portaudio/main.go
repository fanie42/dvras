package main

import (
    "github.com/fanie42/dvras/acquiring/internal/portaudio"
)

func main() {
    repository := nats.NewRepository()
    application := portaudio.NewApplication(repository)
    api := nats.NewAPI(application)
    api.Serve()
}
