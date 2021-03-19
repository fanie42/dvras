package main

import (
    "github.com/fanie42/dvras/internal/http/rest"
    "github.com/fanie42/dvras/internal/portaudio"
)

func main() {
    repo := wav.New()
    app := portaudio.New()
    controller := rest.New(app)

    controller.Run()
}
