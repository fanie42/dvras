package dvras

import (
    "time"

    "github.com/fanie42/dvras/api/protobuf/pb"
    "google.golang.org/protobuf/proto"
)

// Event TODO
type Event interface {
    ID() DeviceID
    Time() time.Time
    Data() []byte
}

// StartedEvent TODO
type StartedEvent struct {
    id DeviceID `json:"id"`
    // Sequence   uint64    `json:"sequence"`
    time       time.Time `json:"time"`
    Annotation string    `json:"annotation"`
}

// ID TODO
func (event *StartedEvent) ID() DeviceID {
    return event.id
}

// Time TODO
func (event *StartedEvent) Time() time.Time {
    return event.time
}

// Data TODO
func (event *StartedEvent) Data() []byte {
    data, _ := proto.Marshal(&pb.StartedEvent{
        Annotation: event.Annotation,
    })
    return data
}

// StoppedEvent TODO
type StoppedEvent struct {
    id DeviceID `json:"id"`
    // Sequence   uint64    `json:"sequence"`
    time       time.Time `json:"time"`
    Annotation string    `json:"annotation"`
}

// ID TODO
func (event *StoppedEvent) ID() DeviceID {
    return event.id
}

// Time TODO
func (event *StoppedEvent) Time() time.Time {
    return event.time
}

// Data TODO
func (event *StoppedEvent) Data() []byte {
    data, _ := proto.Marshal(&pb.StoppedEvent{
        Annotation: event.Annotation,
    })
    return data
}

// DatapointAcquiredEvent TODO
type DatapointAcquiredEvent struct {
    id DeviceID `json:"id"`
    // Sequence uint64    `json:"sequence"`
    time     time.Time `json:"time"`
    Channel1 []int16   `json:"channel_1"`
    Channel2 []int16   `json:"channel_2"`
    PPS      []int16   `json:"pps"`
}

// ID TODO
func (event *DatapointAcquiredEvent) ID() DeviceID {
    return event.id
}

// Time TODO
func (event *DatapointAcquiredEvent) Time() time.Time {
    return event.time
}

// Data TODO
func (event *DatapointAcquiredEvent) Data() []byte {
    ch1 := make([]int32, len(event.Channel1))
    for i, j := range event.Channel1 {
        ch1[i] = int32(j)
    }

    ch2 := make([]int32, len(event.Channel2))
    for i, j := range event.Channel2 {
        ch2[i] = int32(j)
    }

    pps := make([]int32, len(event.PPS))
    for i, j := range event.PPS {
        pps[i] = int32(j)
    }

    data, _ := proto.Marshal(
        &pb.DatapointAcquiredEvent{
            Ch1: ch1,
            Ch2: ch2,
            Pps: pps,
        },
    )

    return data
}
