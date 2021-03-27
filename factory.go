package dvras

import "time"

// Factory TODO
type Factory struct {
    sampleRate uint
    bitDepth   uint
    nsChannel  uint
    ewChannel  uint
    ppsChannel uint
}

// NewFactory TODO
func NewFactory(
    sampleRate uint,
    bitDepth uint,
    nsChannel uint,
    ewChannel uint,
    ppsChannel uint,
) *Factory {
    return &Factory{
        sampleRate: sampleRate,
        bitDepth:   bitDepth,
        nsChannel:  nsChannel,
        ewChannel:  ewChannel,
        ppsChannel: ppsChannel,
    }
}

// New TODO
func (factory *Factory) New(
    timestamp time.Time,
    in [][]int32,
) *Data {
    return &Data{
        Time: timestamp,
        NS:   in[factory.nsChannel],
        EW:   in[factory.ewChannel],
        PPS:  in[factory.ppsChannel],
    }
}
