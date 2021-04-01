package portaudio

import (
    "fmt"
    "time"

    "github.com/fanie42/dvras/internal/acquiring"
    pa "github.com/gordonklaus/portaudio"
)

type gateway struct {
    config Configuration
    // repo   acquiring.Repository
    buffer []int16
    stream *pa.Stream
}

// New TODO
func New(
    configuration Configuration,
    // repository acquiring.Repository,
) (acquiring.SessionGateway, error) {
    gw := &gateway{
        config: configuration,
        // repo:   repository,
        buffer: make([]int16, 44100*2),
    }

    var err error
    gw.stream, err = pa.OpenDefaultStream(
        2, 0,
        float64(44100),
        44100, // config.SampleRate ? 0 in the examples. (maps to default)
        app.callback,
    )
    if err != nil {
        fmt.Printf("failed to open portaudio stream: %v", err)
        return nil
    }

    return gw, nil
}

// Start TODO
func (gw *gateway) Start(command *acquiring.StartCommand) error {
    session, err := app.device.NewSession()
    if err != nil {
        return err
    }

    events, err := session.Start()
    if err != nil {
        return err
    }

    // Handle the events
    app.repo.Save(session, events)

    return nil
}

// Stop TODO
func (gw *gateway) Stop(command *acquiring.StopCommand) error {
    session, err := app.repo.GetSessionByID()
    if err != nil {
        return err
    }

    events, err := session.Stop()
    if err != nil {
        return err
    }

    // Handle the events
    app.repo.Save(session, events)

    return nil
}

func (gw *gateway) callback(in []int16) {
    event := acquiring.DatumAcquiredEvent{
        Time:    time.Now(),
        Samples: in,
    }

    app.dispatch(event)
}
