package portaudio

import (
    "fmt"
    "time"

    dvras "github.com/fanie42/dvras/pkg"
    pa "github.com/gordonklaus/portaudio"
)

// Config TODO
type Config struct {
    SampleRate int            `json:"sample_rate"`
    DeviceID   dvras.DeviceID `json:"device_id"`
}

type service struct {
    config *Config
    device *dvras.Device
    stream *pa.Stream
    buffer [][]int16
}

// New TODO
func New(
    config *Config,
    devices dvras.DeviceRepository,
) dvras.Service {
    var err error
    var device *dvras.Device

    device, err = devices.GetDeviceByID(config.DeviceID)
    if err != nil {
        device = dvras.NewDevice()
        devices.Save(device)

        fmt.Println("could not get device")
    }

    svc := &service{
        config: config,
        device: device,
        buffer: make([][]int16, 2),
    }

    for i := range svc.buffer {
        svc.buffer[i] = make([]int16, config.SampleRate)
    }

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

    // err = svc.stream.Start()
    // if err != nil {
    //     err2 := svc.Stop(
    //         &dvras.StopCommand{
    //             Annotation: "unexpected error",
    //         },
    //     )
    //     if err2 != nil {
    //         fmt.Println(err2)
    //         return svc
    //     }
    //     fmt.Println(err)
    //     return svc
    // }

    // err = svc.device.Start("testing")
    // if err != nil {
    //     fmt.Print("blablabla")
    //     return nil
    // }

    return svc
}

func (svc *service) process(in [][]int16) {
    // device := svc.devices.GetDeviceByID(svc.)
    err := svc.device.AcquireDataPoint(
        time.Now(),
        in[0],
        in[1],
        make([]int16, len(in[0])),
    )
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println(len(in[0]))
}

// Start TODO
func (svc *service) Start(command *dvras.StartCommand) error {
    // device, err := svc.devices.GetDeviceByID(command.DeviceID)
    // if err != nil {
    //     return err
    // }

    err := svc.stream.Start()
    if err != nil {
        err2 := svc.Stop(
            &dvras.StopCommand{
                Annotation: "unexpected error",
            },
        )
        if err2 != nil {
            return err2
        }
        return err
    }

    err = svc.device.Start(command.Annotation)
    if err != nil {
        return err
    }

    // err = svc.repo.Save(svc.device)
    // if err != nil {
    //     // don't remove the changes that were made to the device. What will
    //     // happen if the system reboots when it has a long queue of changes?
    // }

    return err
}

// Stop TODO
func (svc *service) Stop(command *dvras.StopCommand) error {
    // device, err := svc.devices.GetByID(command.DeviceID)
    // if err != nil {
    //     return err
    // }

    err := svc.stream.Stop()
    if err != nil {
        return err
    }

    return svc.device.Stop(command.Annotation)
}

// Close TODO
func (svc *service) Close() error {
    return svc.stream.Close()
}
