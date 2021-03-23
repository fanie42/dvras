package portaudio

import (
    "fmt"
    "time"

    "github.com/fanie42/dvras"
    pa "github.com/gordonklaus/portaudio"
)

type service struct {
    config *Config
    repo   dvras.Repository
    stream *pa.Stream
    buffer [][]int16
}

// New TODO
func New(
    config *Config,
    repository dvras.Repository,
) dvras.Service {
    svc := &service{
        config: config,
        repo:   repository,
        buffer: make([][]int16, 2),
    }

    // device, err := repository.Load(config.DeviceID)
    // if err != nil {
    //     device = dvras.NewDevice()
    //     devices.Save(device)

    //     fmt.Println("could not get device, created a new one")
    // }

    for i := range svc.buffer {
        svc.buffer[i] = make([]int16, config.SampleRate)
    }

    var err error
    svc.stream, err = pa.OpenDefaultStream(
        2,
        0,
        float64(config.SampleRate),
        44100, // config.SampleRate ? This is 0 in the examples...
        svc.process,
    )
    if err != nil {
        fmt.Printf("failed to open portaudio stream: %v", err)
        return nil
    }

    return svc
}

func (svc *service) process(in [][]int16) {
    // Raise event, but don't have to update the aggregate. It won't change
    // the sequence either, I think?
    // err := fmt.Errorf("init loop")
    for {
        device, err := svc.repo.Load(svc.config.DeviceID)
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
        fmt.Println(len(in[0]))

        // Repeat if command fails with "OutOfSequenceError"
        err = svc.repo.Save(device)
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

// Start TODO
func (svc *service) Start(command *dvras.StartCommand) error {
    var device *dvras.Device
    for {
        device, err := svc.repo.Load(svc.config.DeviceID) // Aggregate Query
        if err != nil {
            return err
        }

        // device is the aggregate
        err = device.Start(command.Annotation)
        if err != nil {
            return err
        }

        err = svc.repo.Save(device)
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

    err := svc.stream.Start()
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
func (svc *service) Stop(command *dvras.StopCommand) error {
    for {
        device, err := svc.repo.Load(svc.config.DeviceID) // Aggregate Query
        fmt.Printf("%+v\n", device)
        if err != nil {
            return err
        }

        // device is the aggregate - send to handler with command
        // events saved as part of the updated aggregate
        err = device.Stop(command.Annotation)
        if err != nil {
            return err
        }

        err = svc.repo.Save(device) // Save events list
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

    err := svc.stream.Stop()
    if err != nil {
        return err
    }

    return nil
}

// Close TODO
func (svc *service) Close() error {
    return svc.stream.Close()
}
