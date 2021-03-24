package portaudio

import (
    "fmt"
    "time"

    "github.com/fanie42/dvras"
    pa "github.com/gordonklaus/portaudio"
)

type application struct {
    config *Config
    gw     dvras.Gateway
    stream *pa.Stream
    buffer [][]int16
}

// New TODO
func New(
    config *Config,
    gateway dvras.Gateway,
) dvras.ApplicationService {
    app := &application{
        config: config,
        gw:     gateway,
        buffer: make([][]int16, 2),
    }

    for i := range app.buffer {
        app.buffer[i] = make([]int16, config.SampleRate)
    }

    var err error
    app.stream, err = pa.OpenDefaultStream(
        2,
        0,
        float64(config.SampleRate),
        44100, // config.SampleRate ? This is 0 in the examples...
        app.process,
    )
    if err != nil {
        fmt.Printf("failed to open portaudio stream: %v", err)
        return nil
    }

    return app
}

func (app *application) process(in [][]int16) {
    // Raise event, but don't have to update the aggregate. It won't change
    // the sequence either, I think?
    // err := fmt.Errorf("init loop")
    for {
        device, err := app.gw.Load(app.config.DeviceID)

        if err != nil {
            fmt.Printf("unable to load device: %v", err)
            return
        }

        err = device.AcquireDataPoint(
            time.Now(),
            in[0],
            in[1],
            make([]int16, len(in[0])),
        )
        if err != nil {
            fmt.Printf("command failed: %v", err)
            return
        }

        // Repeat if command fails with "OutOfSequenceError"
        err = app.gw.Save(device)
        if err != nil {
            switch err.(type) {
            case dvras.SequenceConflictError:
                fmt.Println(err)
                continue
            default:
                fmt.Println("unexpected error in portaudio process callback function")
                return
            }
        }

        break
    }

    return
}

// Start TODO - ok... this is all still event sourced. Maybe it would be better
// if the ingress software simply publishes or caches new datapoints? There is
// no validation required for data events to be published, so we don't really
// need event sourcing!?
func (app *application) Start(command *dvras.StartCommand) error {
    var device *dvras.Device
    // This needs to happen every time, so makes sense to somehow abstract this
    // behaviour. Where to?
    for {
        device, err := app.gw.Load(app.config.DeviceID) // Aggregate Query
        if err != nil {
            return err
        }

        // device is the aggregate
        err = device.Start(command.Annotation)
        if err != nil {
            return err
        }

        err = app.gw.Save(device)
        if err != nil {
            switch err.(type) {
            case dvras.SequenceConflictError:
                // The aggregate is out of date and operation will be retried
                // with updated aggregate.
                fmt.Printf("sequence conflict error: %v", err)
                continue
            default:
                // Unknown error saving device. Abort and return error, command
                // fail.
                return err
            }
        }

        break
    }

    // Now actually start the stream. If the start of the physical device fails,
    // We will stop the device automatically with an annotation stating the
    // reason.
    err := app.stream.Start()
    if err != nil {
        err2 := device.Stop("unexpected error while trying to start")
        if err2 != nil {
            return err2
        }
        return err
    }

    return nil
}

// Stop TODO
func (app *application) Stop(command *dvras.StopCommand) error {
    for {
        device, err := app.gw.Load(app.config.DeviceID) // Aggregate Query
        if err != nil {
            return err
        }

        // device is the aggregate - send to handler with command
        // events saved as part of the updated aggregate
        err = device.Stop(command.Annotation)
        if err != nil {
            return err
        }

        err = app.gw.Save(device) // Save events list
        if err != nil {
            switch err.(type) {
            case dvras.SequenceConflictError:
                fmt.Println(err)
                continue
            default:
                return err
            }
        }

        break
    }

    err := app.stream.Stop()
    if err != nil {
        return err
    }

    return nil
}

// Close TODO
func (app *application) Close() error {
    return app.stream.Close()
}
