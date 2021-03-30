package archiving

import (
    "encoding/binary"
    "io"
)

// File TODO
type File struct {
    fn     string
    frames []Frame
}

// NewFile TODO
func NewFile(
    filename string,
) *File {
    return &File{
        fn:     filename,
        frames: []Frame{},
    }
}

func log(err error) {
    log.Printf("error writing: %v", err)
}

func (f *File) wav(file *io.File) {
    numChannels := uint16(len(df.Data))
    sampleRate := uint32(df.SampleRate)
    bitDepth := uint16(df.BitDepth)
    byteRate := uint32(sampleRate * bitDepth * numChannels / 8)
    blockAlign := uint16(bitDepth * numChannels / 8)

    le := binary.LittleEndian

    // RIFF Chunk descriptor
    _, err := file.Write([]byte("RIFF"))   // 4 RIFF
    log(err)                               //
    log(binary.Write(file, le, uint32(0))) // 4 ChunkSize - come back later
    _, err := file.Write([]byte("WAVE"))   // 4 Format
    log(err)                               //

    // fmt Sub-chunk
    _, err := file.Write([]byte("fmt "))     // 4 SubchunkID - fmt
    log(err)                                 //
    log(binary.Write(file, le, uint32(16)))  // 4 SubchunkSize
    log(binary.Write(file, le, uint16(1)))   // 2 AudioFormat
    log(binary.Write(file, le, numChannels)) // 2 NumChannels
    log(binary.Write(file, le, sampleRate))  // 4 SampleRate
    log(binary.Write(file, le, byteRate))    // 4 ByteRate
    log(binary.Write(file, le, uint16(6)))   // 2 BlockAlign
    log(binary.Write(file, le, bitDepth))    // 2 BitsPerSample

    // data Sub-chunk
    _, err := file.Write([]byte("data"))   // 4 SubchunkID - data
    log(err)                               //
    log(binary.Write(file, le, uint32(0))) // 4 SubchunkSize - come back later

    for _, frame := range f.frames {
        for _, datum := range frame.Data {
            log(binary.Write(file, le, []int16{datum.NS, datum.EW, datum.PPS}))
        }
    }
}
