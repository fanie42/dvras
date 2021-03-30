package portaudio

import (
    "fmt"
    "log"

    "github.com/fanie42/dvras"
    pa "github.com/gordonklaus/portaudio"
)

type application struct {
    config *Config
    repo   dvras.Repository
    stream *pa.Stream
    buffer []int16
}

// New TODO
func New(
    config *Config,
    repository dvras.Repository,
) dvras.ApplicationService {
    app := &application{
        config: config,
        repo:   repository,
        buffer: []int16{},
    }

    for i := range app.buffer {
        app.buffer[i] = make([]int16, config.SampleRate)
    }

    var err error
    app.stream, err = pa.OpenDefaultStream(
        2, 0,
        float64(44100),
        44100, // config.SampleRate ? This is 0 in the examples...
        app.callback,
    )
    if err != nil {
        fmt.Printf("failed to open portaudio stream: %v", err)
        return nil
    }

    return app
}

func (app *application) callback(in [][]int16) {
    device, err := app.repo.Load(app.config.DeviceID)
    if err != nil {
        fmt.Printf("unable to load device: %v", err)
        return
    }

    err = device.AcquireData(
        in[0],
        in[1],
        make([]int16, len(in[0])),
    )
    if err != nil {
        fmt.Printf("command failed: %v", err)
        return
    }

    err = app.repo.Save(device)
    if err != nil {
        fmt.Println(err)
    }
}

// Start TODO - ok... this is all still event sourced. Maybe it would be better
// if the ingress software simply publishes or caches new datapoints? There is
// no validation required for data events to be published, so we don't really
// need event sourcing!?
func (app *application) Start(
    annotation string,
) error {
    device, err := app.repo.Load(app.config.DeviceID)
    if err != nil {
        return err
    }

    command = device.Start(annotation, app.stream.Start)
    events, err := command.Execute()
    if err != nil {
        return err
    }

    err = app.repo.Save(device, events)
    if err != nil {
        log.Printf("failed to persist started event: %v", err)
    }

    return nil
}

// Stop TODO
func (app *application) Stop(
    annotation string,
) error {
    device, err := app.repo.Load(app.config.DeviceID)
    if err != nil {
        return err
    }

    err = device.Stop(
        annotation,
        app.stream.Stop,
    )
    if err != nil {
        return err
    }

    err = app.repo.Save(device)
    if err != nil {
        log.Printf("failed to persist stopped event: %v", err)
        return err
    }

    return nil
}

// Close TODO
func (app *application) Close() error {
    return app.stream.Close()
}
