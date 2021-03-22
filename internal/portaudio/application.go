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
    err := fmt.Errorf("init loop")
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
        if err != dvras.OutOfSequenceError {
            break
        }
    }

    return
}

// Start TODO
func (svc *service) Start(command *dvras.StartCommand) error {
    device, err := svc.repo.Load(svc.config.DeviceID)
    if err != nil {
        return err
    }

    err = device.Start(command.Annotation)
    if err != nil {
        return err
    }

    err = svc.stream.Start()
    if err != nil {
        err2 := device.Stop("unexpected error while trying to start")
        if err2 != nil {
            return err2
        }
        return err
    }

    svc.repo.Save(device)

    // err = svc.repo.Save(svc.device)
    // if err != nil {
    //     // don't remove the changes that were made to the device. What will
    //     // happen if the system reboots when it has a long queue of changes?
    // }

    return err
}

// Stop TODO
func (svc *service) Stop(command *dvras.StopCommand) error {
    device, err := svc.repo.Load(svc.config.DeviceID)
    if err != nil {
        return err
    }

    err = device.Stop(command.Annotation)
    if err != nil {
        return err
    }

    err = svc.stream.Stop()
    if err != nil {
        err2 := device.Start("unexpected error while trying to stop")
        if err2 != nil {
            return err2
        }
        return err
    }

    svc.repo.Save(device)

    return err
}

// Close TODO
func (svc *service) Close() error {
    return svc.stream.Close()
}
