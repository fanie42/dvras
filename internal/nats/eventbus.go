package nats

import (
    "fmt"
    "time"

    "github.com/fanie42/dvras"
    "github.com/fanie42/dvras/api/protobuf/pb"
    "github.com/fanie42/dvras/vendor/google.golang.org/protobuf/proto"
    "github.com/google/uuid"
)

// eventbus TODO
type eventbus struct {
    nc *natsio.Conn
}

// New TODO
func New(
    conn *natsio.Conn,
) monitoring.EventBus {
    eb := &eventbus{
        nc: conn,
    }

    return eb
}

// OnStarted TODO
func (eb *eventbus) OnStarted(
    handle func(*dvras.StartedEvent) error,
) error {
    sub, err := eb.nc.Subscribe(
        "dvras.device.started",
        func(msg *natsio.Msg) {
            event := &pb.StartedEvent{}
            if err := proto.Unmarshal(msg.Data, event); err != nil {
                fmt.Printf("failed to parse event: %v", err)
            }
            id, _ := uuid.Parse(event.DeviceId.Uuid)
            handle(&dvras.StartedEvent{
                DeviceID:   id,
                Annotation: event.Annotation,
            })
        },
    )
    if err != nil {
        return fmt.Errorf("failed to subscribe: %v", err)
    }

    eb.subs = append(eb.subs, sub)

    return nil
}

// OnStopped TODO
func (eb *eventbus) OnStopped(
    handle func(*dvras.StoppedEvent) error,
) error {
    sub, err := eb.nc.Subscribe(
        "dvras.device.stopped",
        func(msg *natsio.Msg) {
            event := &pb.StoppedEvent{}
            if err := proto.Unmarshal(msg.Data, event); err != nil {
                fmt.Printf("failed to parse event: %v", err)
            }
            id, _ := uuid.Parse(event.DeviceId.Uuid)
            handle(&dvras.StoppedEvent{
                DeviceID:   id,
                Annotation: event.Annotation,
            })
        },
    )
    if err != nil {
        return fmt.Errorf("failed to subscribe: %v", err)
    }

    eb.subs = append(eb.subs, sub)

    return nil
}

// OnDatapointAcquired TODO
func (eb *eventbus) OnDatapointAcquired(
    handle func(*dvras.DatapointAcquiredEvent) error,
) error {
    sub, err := eb.nc.Subscribe(
        "dvras.device.datapoint.acquired",
        func(msg *natsio.Msg) {
            event := &pb.DatapointAcquiredEvent{}
            if err := proto.Unmarshal(msg.Data, event); err != nil {
                fmt.Printf("failed to parse event: %v", err)
            }
            id, _ := uuid.Parse(event.DeviceId.Uuid)
            ch1 := make([]int16, len(event.Data))
            ch2 := make([]int16, len(event.Data))
            pps := make([]int16, len(event.Data))
            for i, d := range event.Data {
                ch1[i] = int16(d.Ch1)
                ch2[i] = int16(d.Ch2)
                pps[i] = int16(d.Pps)
            }
            handle(&dvras.DatapointAcquiredEvent{
                Timestamp: time.Unix(event.Timestamp.Seconds, event.Timestamp.Nanoseconds),
                DeviceID:  id,
                Channel1:  ch1,
                Channel2:  ch2,
                PPS:       pps,
            })
        },
    )
    if err != nil {
        return fmt.Errorf("failed to subscribe: %v", err)
    }

    eb.subs = append(eb.subs, sub)

    return nil
}

func (app *Application) onStarted(
    msg *natsio.Msg,
) {
    event := dvras.StartedEvent{}
    err := event.Unmarshal(msg.Data)
    if err != nil {
        fmt.Println(err)
    }
    app.proj.OnStarted(event)
}

func (app *Application) onStopped(
    msg *natsio.Msg,
) {
    event := dvras.StoppedEvent{}
    err := event.Unmarshal(msg.Data)
    if err != nil {
        fmt.Println(err)
    }
    app.proj.OnStopped(event)
}

func (app *Application) onDatapointAcquired(
    msg *natsio.Msg,
) {
    event := dvras.DatapointAcquiredEvent{}
    err := event.Unmarshal(msg.Data)
    if err != nil {
        fmt.Println(err)
    }
    app.proj.OnDatapointAcquired(event)
}
