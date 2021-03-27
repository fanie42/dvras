package portaudio

import (
    "fmt"
    "log"
    "time"

    "github.com/fanie42/dvras"
    pa "github.com/gordonklaus/portaudio"
)

type application struct {
    config *Configuration
    repo   dvras.Repository
    stream *pa.Stream
    buffer [][]int32
}

// New TODO
func New(
    configuration *Configuration,
    repository dvras.Repository,
) dvras.ApplicationService {
    app := &application{
        config: configuration,
        repo:   repository,
        buffer: make([][]int32, 2),
    }

    for i := range app.buffer {
        app.buffer[i] = make([]int32, 44100)
    }

    var err error
    app.stream, err = pa.OpenDefaultStream(
        2, 0,
        float64(44100),
        44100,
        app.callback,
    )
    if err != nil {
        log.Fatalf("failed to open portaudio stream: %v", err)
        return nil
    }

    return app
}

// Start TODO
func (app *application) Start() error {
    err := app.stream.Start()
    if err != nil {
        return err
    }

    return nil
}

// Stop TODO
func (app *application) Stop() error {
    err := app.stream.Stop()
    if err != nil {
        return err
    }

    return nil
}

func (app *application) callback(buffer [][]int32) {
    timestamp := time.Now()

    data := dvras.New(timestamp)

    err := data.Acquire(
        buffer[app.config.NSchannel],
        buffer[app.config.EWchannel],
        make([]int32, len(buffer[0])),
    )
    if err != nil {
        log.Printf("acquisition failed: %v", err)
        return
    }

    err = app.repo.Save(data)
    if err != nil {
        fmt.Printf("failed to save data at time: %v", timestamp)
        return
    }

    fmt.Println(len(buffer[0]))

    return
}

// Close TODO
func (app *application) Close() error {
    return app.stream.Close()
}
