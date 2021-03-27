package dvras

import (
    "encoding/json"
    "log"
)

// Event TODO
type Event interface {
    Data() []byte
}

// // StartedEvent TODO
// type StartedEvent struct {
//     Annotation string
// }

// // StoppedEvent TODO
// type StoppedEvent struct {
//     Annotation string
// }

// AcquiredEvent TODO
type AcquiredEvent struct {
    NS  []int32 `json:"ns"`
    EW  []int32 `json:"ew"`
    PPS []int32 `json:"pps"`
}

// Data TODO
func (event *AcquiredEvent) Data() []byte {
    b, err := json.Marshal(event)
    if err != nil {
        log.Printf("failed to marshal event: %v", err)
    }

    return b
}
