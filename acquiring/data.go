package acquiring

import "time"

// Datum TODO
type Datum struct {
    SessionID string    `json:"session_id"`
    Time      time.Time `json:"time"`
    NS        []int16   `json:"ns"`
    EW        []int16   `json:"ew"`
    PPS       []int16   `json:"pps"`
}
