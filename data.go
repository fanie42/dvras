package dvras

import "time"

// Data TODO
type Data struct {
    Time time.Time
    NS   []int16
    EW   []int16
    PPS  []int16
}
