package archiving

import (
    "log"
    "os"
    "time"
)

// Archiver TODO
type Archiver interface {
    Archive(*Datum)
}

type archiver struct {
    files  []File
    layout string
    sr     uint32
    bd     uint16
}

// Frame TODO
type Frame struct {
    Time time.Time
    Data []Datum
}

// Datum TODO
type Datum struct {
    NS  int16
    EW  int16
    PPS int16
}

// New TODO
func New(
    path, base string,
    sampleRate uint32,
    bitDepth uint16,
) Archiver {
    return &archiver{
        files:  []File{},
        layout: path.Join(path, base),
        sr:     sampleRate,
        bd:     bitDepth,
    }
}

// Archive TODO
func (a *archiver) Archive(f *Frame) {
    if f.Time-a.now > time.Duration(time.Minute) {
        // Start new cache with this frame and write out previous cache
        if len(a.buffer > 0) {
            temp := []Frame{}
            n := copy(temp, a.buffer)
            go a.wav(temp)
        }
    }

    if len(a.cache) < 1 {
        a.cache = append(a.cache, f)
        a.now = f.Time
    }

    if !ok {
        frames = []Frame{}
        frames = append(frames, &Frame{
            Time: d.Time,
            NS:   d.NS,
            EW:   d.EW,
            PPS:  d.PPS,
        })
    }

    a.cache[d.DeviceID] = frames

    data = append(data, &Frame{
        Time: d.Time,
        NS:   d.NS,
        EW:   d.EW,
        PPS:  d.PPS,
    })
}

// Archive TODO
func (f *File) wav() {
    file, err := os.OpenFile(
        data.Time.Format(f.fn),
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0644,
    )
    if err != nil {
        return err
    }

    _, _ := file.Write(data)

    if err := file.Close(); err != nil {
        log.Printf("failed to close file: ", err)
        return err
    }

    return nil
}
